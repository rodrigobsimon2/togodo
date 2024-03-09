package service

import (
	"testing"

	"github.com/rodrigobsimon/togodo/internal/usecase"
)

func TestSaveTask(t *testing.T) {
	pi := PersistentInterfaceMock{}
	task := usecase.CreateTask("Some task", usecase.IMPORTANT_URGENT)
	SaveTask(task, pi)
}