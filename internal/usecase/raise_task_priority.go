package usecase

func (t *Task) RaiseTaskPriority() {
	if t.Priority != IMPORTANT_URGENT {
		t.Priority = t.Priority+1
	}
}
