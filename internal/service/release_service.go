package service

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func (s *System) Complete(id string, decision domain.CompletionDecision) error {
	if !decision.Allowed {
		return domain.ErrUnsafeComplete
	}
	return s.Registry.Mutate(id, func(p *domain.RepairPlan) error {
		if p.State == domain.RepairAborted {
			return domain.ErrUnsafeComplete
		}
		if p.State == domain.RepairScheduled {
			if err := domain.Transition(p, domain.RepairExecuting); err != nil {
				return err
			}
		}
		if p.State == domain.RepairExecuting {
			return domain.Transition(p, domain.RepairCompleted)
		}
		return domain.ErrInvalidState
	})
}
func (s *System) CompleteWithEvents(id string, decision domain.CompletionDecision) error {
	if err := s.Complete(id, decision); err != nil {
		return err
	}
	s.Events.Append("completed:" + id)
	return nil
}
