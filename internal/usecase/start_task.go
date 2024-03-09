package usecase

func (t *Task) StartTask() {
	if t.Status == TODO {
		t.Status = DOING
	}
}