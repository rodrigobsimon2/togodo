package usecase

type Task struct {
	Description string
	Status      int
	Priority    int
}

const (
	TODO = iota
	DOING
	DONE
)

const (
	NOT_IMPORTANT_NOT_URGENT = iota
	NOT_IMPORTANT_URGENT
	IMPORTANT_NOT_URGENT
	IMPORTANT_URGENT
)

func CreateTask(description string, priority int) Task {
	return Task{Description: description, Status: TODO, Priority: priority}
}
