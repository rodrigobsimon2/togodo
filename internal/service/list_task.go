package service

import "github.com/rodrigobsimon/togodo/internal/usecase"

func ListTask(pi PersistentInterface) []usecase.Task {
	ts, err := pi.ListTask()
	if err != nil {
		panic(err)
	}
	return ts
}