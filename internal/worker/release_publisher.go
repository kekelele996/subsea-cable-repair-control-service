package worker

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func PublishComplete(plan domain.RepairPlan, decision domain.CompletionDecision, emit func(string)) error {
	if !decision.Allowed || plan.State == domain.RepairAborted || plan.State == domain.RepairCompleted {
		return domain.ErrUnsafeComplete
	}
	if plan.State != domain.RepairExecuting {
		return domain.ErrInvalidState
	}
	emit("completed:" + plan.ID)
	return nil
}
