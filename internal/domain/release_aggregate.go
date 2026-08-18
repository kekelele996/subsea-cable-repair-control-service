package domain

type CompleteAggregate struct {
	State    RepairState
	Revision int
}

func (a *CompleteAggregate) ApplyComplete(expectedRevision int) error {
	if a.Revision != expectedRevision || a.State == RepairAborted || a.State == RepairCompleted {
		return ErrInvalidState
	}
	if a.State != RepairExecuting {
		return ErrUnsafeComplete
	}
	a.State = RepairCompleted
	a.Revision++
	return nil
}
