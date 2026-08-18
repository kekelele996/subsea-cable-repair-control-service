package service

import (
	"context"
	"fmt"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"time"
)

func (s *System) CreatePlan(id, cable string, spans []string) domain.RepairPlan {
	p := domain.RepairPlan{ID: id, CableID: cable, Spans: domain.CloneStrings(spans), State: domain.RepairDraft, CreatedAt: s.now()}
	s.Registry.Put(p)
	return p
}
func (s *System) QueuePlan(ctx context.Context, id, owner string) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("queue: %v", err)
	}
	plan, ok := s.Registry.Get(id)
	if !ok {
		return domain.ErrInvalidState
	}
	if err := s.Leases.Acquire(plan.CableID, owner, time.Minute); err != nil {
		return fmt.Errorf("queue lease: %v", err)
	}
	if err := s.Registry.Mutate(id, func(plan *domain.RepairPlan) error { return domain.Transition(plan, domain.RepairScheduled) }); err != nil {
		s.Leases.Complete(plan.CableID, owner)
		return err
	}
	s.Events.Append("queued:" + id)
	return nil
}
func (s *System) RejectPlan(id string) error {
	return s.Registry.Mutate(id, func(p *domain.RepairPlan) error { return domain.Transition(p, domain.RepairAborted) })
}
