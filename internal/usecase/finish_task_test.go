package usecase

import "testing"

func TestFinishTask(t *testing.T) {
	task := CreateTask("Walk Dog", IMPORTANT_URGENT)
	task.StartTask()
	task.FinishTask()
	if task.Status != DONE {
		t.Errorf("expected %d, got %d", DONE, task.Status)
	}
}