package task

import (
	"crypto/aes"
	"crypto/cipher"
)

type Service struct {
	tasks map[uint64]Task

}

func (s *Service) AddTask(task Task) {
	s.tasks[task.Id] = task
}

func (s *Service) decrypt(task *Task, key []byte) *Result {
	decryptedData := make([]byte, 0)

	if task.AlgorithmUsed == AES {
		c, err := aes.NewCipher(key)
		if err != nil {
			return nil
		}

		gcm, err := cipher.NewGCM(c)
		if err != nil {
			return nil
		}

		nonceSize := gcm.NonceSize()
		if len(task.Target) < nonceSize {
			return nil
		}

		nonce, ciphertext := task.Target[:nonceSize], task.Target[nonceSize:]
		decryptedData, err = gcm.Open(nil, nonce, ciphertext, nil)

		if err != nil {
			return nil
		}
	}

	j := 0
	for i := 0; i < len(decryptedData); i++ {
		cbyte := decryptedData[i]

		if task.PartialData[j] == cbyte {
			j++
			if j == len(task.PartialData) {
				return &Result{
					Id:     task.Id,
					Output: key,
				}
			}
		} else {
			j = 0
			if i != 0 && j != 0 {
				i--
			}
		}
	}

	return nil
}

func (s *Service) BruteForce(task Task) *Result {
	payload := task.PayloadData
	alphabetSize := len(payload.Alphabet)
	alphabet := payload.Alphabet
	key := payload.Start

	normalizedKey := make([]byte, 0)
	for i := 0; i < len(key); i++ {
		keyIndex := 0
		for keyIndex < alphabetSize {
			if key[i] == alphabet[keyIndex] {
				break
			}

			keyIndex++
		}

		normalizedKey = append(normalizedKey, byte(keyIndex))
	}

	for i := uint64(0); i < payload.Count; i++ {
		// fmt.Println("Trying password: ", key)
		result := s.decrypt(&task, append(payload.Prefix, key...))
		if result != nil {
			return result
		}

		n := len(key)
		prependNewDigit := true

		for j := n - 1; j >= 0; j-- {
			if normalizedKey[j] == byte(alphabetSize - 1) {
				normalizedKey[j] = 0
				key[j] = alphabet[0]
			} else {
				normalizedKey[j]++
				key[j] = alphabet[normalizedKey[j]]
				prependNewDigit = false
				break
			}
		}

		if prependNewDigit {
			normalizedKey = append([]byte{0}, normalizedKey...)
			key = append([]byte{alphabet[normalizedKey[0]]}, key...)
		}
	}

	return nil
}

func NewTaskService() Service {
	return Service{
		tasks: make(map[uint64]Task),
	}
}
