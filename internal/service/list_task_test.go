package service

import (
	"testing"

	"github.com/rodrigobsimon/togodo/internal/usecase"
)

type PersistentInterfaceMock struct{}

func (pi PersistentInterfaceMock) ListTask() ([]usecase.Task, error) {
	return []usecase.Task{
		usecase.CreateTask("Some task", usecase.IMPORTANT_URGENT),
		usecase.CreateTask("Another task", usecase.NOT_IMPORTANT_NOT_URGENT),
	}, nil
}

func (pi PersistentInterfaceMock) SaveTask(t usecase.Task) error {
	return nil
}

func TestListTask(t *testing.T) {
	pi := PersistentInterfaceMock{}
	ts := ListTask(pi)
	expected := 2
	if len(ts) != expected {
		t.Errorf("expected %d, got %d", expected, len(ts))
	}
}