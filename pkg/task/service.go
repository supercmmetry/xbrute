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

func (s *Service) Decrypt(task Task, key []byte) *Result {
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
				return &Result {
					Id: task.Id,
					Output: key,
				}
			}
		} else {
			j = 0
		}
	}

	return nil
}

func NewTaskService() Service {
	return Service{
		tasks: make(map[uint64]Task),
	}
}
