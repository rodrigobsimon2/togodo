package service

import "github.com/rodrigobsimon/togodo/internal/usecase"

type PersistentInterface interface {
	SaveTask(t usecase.Task) error
	ListTask() ([]usecase.Task, error)
}

func SaveTask(t usecase.Task, pi PersistentInterface)  {
	err := pi.SaveTask(t)
	if err != nil {
		panic(err)
	}
}