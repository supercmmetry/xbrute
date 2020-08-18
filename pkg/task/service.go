package task

type Service struct {
	tasks []Task
}

func NewTaskService() Service {
	return Service{
		tasks: make([]Task, 0),
	}
}
