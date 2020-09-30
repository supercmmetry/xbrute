package task

type Payload struct {
	Prefix   []byte `json:"prefix"`
	Start    []byte `json:"start"`
	Count    uint64 `json:"count"`
	Alphabet []byte `json:"alphabet"`
}

type Algorithm uint64

type Task struct {
	Id            uint64  `json:"id"`
	PayloadData   Payload `json:"payload"`
	Target        []byte  `json:"target"`
	PartialData   []byte  `json:"partial_data"`
	AlgorithmUsed string  `json:"algorithm"`
	AttackCount   uint64  `json:"attack_count"`
	Solution      []byte  `json:"solution"`
}

type Result struct {
	Id     uint64 `json:"id"`
	Output []byte `json:"output"`
}
