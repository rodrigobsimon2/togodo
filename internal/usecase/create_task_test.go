package usecase

import "testing"

func TestCreateTask(t *testing.T) {
	expectedDescription := "Walk dog"
	task := CreateTask("Walk dog", IMPORTANT_URGENT)
	if task.Description != expectedDescription {
		t.Errorf("expected %s, task %s", expectedDescription, task.Description)
	}

	if task.Status != TODO {
		t.Errorf("expected %d, task %d", TODO, task.Status)
	}

	if task.Priority != IMPORTANT_URGENT {
		t.Errorf("expected %d, task %d", IMPORTANT_URGENT, task.Priority)
	}
}
