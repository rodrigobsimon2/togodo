package usecase

import "testing"

func TestDecreaseTaskPriority(t *testing.T){
	task := CreateTask("Walk Dog", IMPORTANT_URGENT)
	task.DecreaseTaskPriority()
	if task.Priority != IMPORTANT_NOT_URGENT {
		t.Errorf("expected %d, got %d", task.Priority, IMPORTANT_NOT_URGENT)
	}

	task.DecreaseTaskPriority()
	task.DecreaseTaskPriority()
	task.DecreaseTaskPriority()

	if task.Priority != NOT_IMPORTANT_NOT_URGENT {
		t.Errorf("expected %d, got %d", task.Priority, NOT_IMPORTANT_NOT_URGENT)
	}
}