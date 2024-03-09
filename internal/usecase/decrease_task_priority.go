package usecase

func (t *Task) DecreaseTaskPriority()  {
	if t.Priority != NOT_IMPORTANT_NOT_URGENT {
		t.Priority = t.Priority-1
	}
}