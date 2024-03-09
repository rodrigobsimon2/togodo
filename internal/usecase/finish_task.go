package usecase

func (t *Task) FinishTask() {
	if t.Status == DOING {
		t.Status = DONE
	}
}