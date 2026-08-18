package service_test

import (
	"errors"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/service"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/store"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/worker"
	"testing"
)

type nilSafeSafety struct{}

func (p *nilSafeSafety) Evaluate(domain.RepairPlan) ([]domain.Hazard, error) {
	if p == nil {
		return []domain.Hazard{{Code: "nil-provider", Severity: "low"}}, nil
	}
	return nil, nil
}
func TestTypedNilSafetyRejectedAtEveryBoundary(t *testing.T) {
	var provider *nilSafeSafety
	var iface domain.SafetyProvider = provider
	if (domain.ProviderHandle{Provider: iface}).Available() {
		t.Fatal("domain handle reported a typed-nil provider as available")
	}
	registry := store.NewSafetyRegistry()
	if registry.Register("complete", iface) || registry.Has("complete") {
		t.Fatal("registry accepted a typed-nil provider")
	}
	if _, err := service.EvaluateProvider(domain.RepairPlan{ID: "p4"}, iface); !errors.Is(err, domain.ErrSafetyUnavailable) {
		t.Fatalf("service evaluated typed-nil provider: %v", err)
	}
	if err := worker.AcceptSafetyJob(worker.SafetyJob{Provider: iface}); !errors.Is(err, domain.ErrSafetyUnavailable) {
		t.Fatalf("worker accepted typed-nil provider: %v", err)
	}
}
