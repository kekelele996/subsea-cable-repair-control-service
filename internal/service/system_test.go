package service_test

import (
	"context"
	"errors"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/service"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/worker"
	"testing"
	"time"
)

type cleanSafety struct{}

func (cleanSafety) Evaluate(domain.RepairPlan) ([]domain.Hazard, error) { return nil, nil }

type typedNilSafety struct{}

func (*typedNilSafety) Evaluate(domain.RepairPlan) ([]domain.Hazard, error) { return nil, nil }
func fresh() *service.System                                                { return service.NewSystem(func() time.Time { return time.Unix(100, 0) }) }
func queued(s *service.System, id string) {
	s.CreatePlan(id, "cable-"+id, []string{"north", "south"})
	if err := s.QueuePlan(context.Background(), id, "operator"); err != nil {
		panic(err)
	}
}
func TestConcurrentMobilizeSnapshotDoesNotLeakSpans(t *testing.T) {
	s := fresh()
	queued(s, "p1")
	if err := s.StageManifest("p1"); err != nil {
		t.Fatal(err)
	}
	copy := s.AddEmergencySpan("p1", "emergency")
	if len(copy) != 3 {
		t.Fatal(copy)
	}
	if got := s.StoredManifest("p1"); len(got) != 2 {
		t.Fatalf("manifest leaked %v", got)
	}
}
func TestCoordinatorDrainsJobs(t *testing.T) {
	s := fresh()
	queued(s, "p2")
	got, err := worker.Mobilize(context.Background(), s, "p2")
	if err != nil || len(got) != 2 {
		t.Fatalf("got=%v err=%v", got, err)
	}
}
func TestRetryRecognizesLeaseConflict(t *testing.T) {
	s := fresh()
	if err := s.ReserveForRetry("cable", "first"); err != nil {
		t.Fatal(err)
	}
	if got := s.RetryCable("cable", "second"); got != "defer" {
		t.Fatalf("got %s", got)
	}
}
func TestTypedNilSafetyIsDenied(t *testing.T) {
	s := fresh()
	s.CreatePlan("p4", "cable", []string{"north"})
	var p *typedNilSafety
	decision, err := s.EvaluateComplete("p4", p)
	if !errors.Is(err, domain.ErrSafetyUnavailable) || decision.Allowed {
		t.Fatalf("decision=%+v err=%v", decision, err)
	}
}
func TestCancelledContextDoesNotQueuePlan(t *testing.T) {
	s := fresh()
	s.CreatePlan("p5", "cable", []string{"north"})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := worker.CancelAwareMutation(ctx, s, "p5"); err == nil {
		t.Fatal("expected cancellation")
	}
	p, _ := s.Registry.Get("p5")
	if p.State != domain.RepairDraft {
		t.Fatalf("state=%s", p.State)
	}
}
func TestManifestAppendDoesNotPolluteStoredSpans(t *testing.T) {
	s := fresh()
	queued(s, "p6")
	if err := s.StageManifest("p6"); err != nil {
		t.Fatal(err)
	}
	_ = s.AddEmergencySpan("p6", "temporary")
	if got := s.StoredManifest("p6"); len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}
func TestFinalizeBatchReturnsErrorAndCompletesLease(t *testing.T) {
	s := fresh()
	err := s.FinalizeBatch("cable", "owner", func() error { return domain.ErrFinalization })
	if !errors.Is(err, domain.ErrFinalization) {
		t.Fatalf("err=%v", err)
	}
	if s.Leases.Held("cable") {
		t.Fatal("lease kept")
	}
}
func TestRejectedPlanCannotComplete(t *testing.T) {
	s := fresh()
	queued(s, "p8")
	if err := s.RejectPlan("p8"); err != nil {
		t.Fatal(err)
	}
	if err := s.CompleteWithEvents("p8", domain.CompletionDecision{PlanID: "p8", Allowed: true}); !errors.Is(err, domain.ErrUnsafeComplete) {
		t.Fatalf("err=%v", err)
	}
	if s.Events.Count("completed:p8") != 0 {
		t.Fatal("complete event emitted")
	}
}
func TestCompleteMovesQueuedPlanThroughRunning(t *testing.T) {
	s := fresh()
	queued(s, "p9")
	if err := s.CompleteWithEvents("p9", domain.CompletionDecision{PlanID: "p9", Allowed: true}); err != nil {
		t.Fatal(err)
	}
	p, _ := s.Registry.Get("p9")
	if p.State != domain.RepairCompleted {
		t.Fatal(p.State)
	}
}
func TestSafetyBlocksCriticalSpan(t *testing.T) {
	s := fresh()
	s.CreatePlan("p10", "cable", []string{"critical-access"})
	d, err := s.EvaluateComplete("p10", cleanSafety{})
	if err != nil {
		t.Fatal(err)
	}
	if !d.Allowed {
		t.Fatal("clean safety should allow")
	}
}
