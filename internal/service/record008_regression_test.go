package service_test

import (
	"errors"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/service"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/store"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/worker"
	"testing"
)

func TestRejectedCompleteIsFencedAtEveryBoundary(t *testing.T) {
	aggregate := domain.CompleteAggregate{State: domain.RepairAborted, Revision: 7}
	if err := aggregate.ApplyComplete(7); !errors.Is(err, domain.ErrInvalidState) || aggregate.State != domain.RepairAborted {
		t.Fatalf("aggregate completed terminal state: %+v err=%v", aggregate, err)
	}
	plan := domain.RepairPlan{ID: "p8", State: domain.RepairAborted, Revision: 7}
	repo := store.NewCompleteRepository(plan)
	if err := repo.Commit(7); !errors.Is(err, domain.ErrInvalidState) || repo.Plan.State != domain.RepairAborted {
		t.Fatalf("repository completed rejected plan: %+v err=%v", repo.Plan, err)
	}
	repo2 := store.NewCompleteRepository(plan)
	coordinator := &service.CompleteCoordinator{}
	if err := coordinator.Complete(repo2, 7); !errors.Is(err, domain.ErrInvalidState) || len(coordinator.Events) != 0 {
		t.Fatalf("coordinator emitted before rejected commit: events=%v err=%v", coordinator.Events, err)
	}
	published := []string{}
	err := worker.PublishComplete(plan, domain.CompletionDecision{PlanID: "p8", Allowed: true}, func(event string) { published = append(published, event) })
	if !errors.Is(err, domain.ErrUnsafeComplete) || len(published) != 0 {
		t.Fatalf("publisher emitted rejected complete: events=%v err=%v", published, err)
	}
}
