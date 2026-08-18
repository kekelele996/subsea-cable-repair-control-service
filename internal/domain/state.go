package domain

var allowedTransitions = map[RepairState]map[RepairState]bool{
	RepairDraft:     {RepairScheduled: true, RepairAborted: true},
	RepairScheduled: {RepairExecuting: true, RepairAborted: true},
	RepairExecuting: {RepairAborted: true, RepairCompleted: true},
	RepairAborted:   {},
	RepairCompleted: {},
}

func CanTransition(from, to RepairState) bool { return allowedTransitions[from][to] }
func Transition(plan *RepairPlan, to RepairState) error {
	if !CanTransition(plan.State, to) {
		return ErrInvalidState
	}
	plan.State = to
	plan.Revision++
	return nil
}
