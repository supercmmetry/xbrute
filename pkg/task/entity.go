package task

type Payload struct {
	Start []byte
	Count uint64
}

type Task struct {
	Id          uint64
	PayloadData Payload
	Target      []byte
	PartialData []byte
}
