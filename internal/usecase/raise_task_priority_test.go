package usecase

import "testing"

func TestRaiseTaskPriority(t *testing.T) {
	task := CreateTask("Walk dog", NOT_IMPORTANT_NOT_URGENT)
	task.RaiseTaskPriority()
	if task.Priority != NOT_IMPORTANT_URGENT {
		t.Errorf("expected %d, got %d", NOT_IMPORTANT_URGENT, task.Priority)
	}

	task.RaiseTaskPriority()
	task.RaiseTaskPriority()
	task.RaiseTaskPriority()

	if task.Priority != IMPORTANT_URGENT {
		t.Errorf("expected %d, got %d", IMPORTANT_URGENT, task.Priority)
	}
}