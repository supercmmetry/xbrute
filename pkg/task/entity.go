package task

type Payload struct {
	Start    []byte `json:"start"`
	Count    uint64 `json:"count"`
	Alphabet []byte `json:"alphabet"`
}

type Algorithm uint64

const (
	AES Algorithm = iota
	RSA
)

type Task struct {
	Id            uint64    `json:"id"`
	PayloadData   Payload   `json:"payload"`
	Target        []byte    `json:"target"`
	PartialData   []byte    `json:"partial_data"`
	AlgorithmUsed Algorithm `json:"algorithm"`
}

type Result struct {
	Id     uint64 `json:"id"`
	Output []byte `json:"output"`
}
