package usecase

import "testing"

func TestStartTask(t *testing.T) {
	task := CreateTask("Walk Dog", IMPORTANT_URGENT)
	task.StartTask()
	if task.Status != DOING {
		t.Errorf("expected %d, got %d", DOING, task.Status)
	}
}