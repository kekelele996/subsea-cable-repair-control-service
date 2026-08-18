package domain

type CompleteAggregate struct {
	State    RepairState
	Revision int
}

func (a *CompleteAggregate) ApplyComplete(expectedRevision int) error {
	a.State = RepairCompleted
	a.Revision++
	return nil
}
