package worker

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func PublishComplete(plan domain.RepairPlan, decision domain.CompletionDecision, emit func(string)) error {
	if !decision.Allowed {
		return domain.ErrUnsafeComplete
	}
	emit("completed:" + plan.ID)
	return nil
}
