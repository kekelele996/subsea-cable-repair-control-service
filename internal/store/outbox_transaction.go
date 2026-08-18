package store

type OutboxTransaction struct {
	RolledBack bool
	Committed  bool
}

func (t *OutboxTransaction) Finish(commit func() error) (err error) {
	defer func() {
		if err != nil {
			t.RolledBack = true
			return
		}
		t.Committed = true
	}()
	return commit()
}
