package workercheck

import (
	"context"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/service"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/store"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/worker"
)

func TestCoordinatorKeepsAllThreeRepairSpans(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	gate := domain.NewCompletionGate(3)
	completeFollowers := make(chan struct{})
	firstCompleted := make(chan struct{})
	var completeOnce sync.Once
	complete := func() { completeOnce.Do(func() { close(completeFollowers) }) }
	defer complete()
	var participants sync.WaitGroup
	participants.Add(3)
	go func() {
		defer participants.Done()
		gate.Complete()
		close(firstCompleted)
	}()
	go func() {
		defer participants.Done()
		<-completeFollowers
		gate.Complete()
	}()
	go func() {
		defer participants.Done()
		<-completeFollowers
		gate.Complete()
	}()
	<-firstCompleted
	short, cancelShort := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelShort()
	if gate.Wait(short) == nil {
		t.Fatal("domain completion gate opened before all spans finished")
	}
	complete()
	participants.Wait()
	if err := gate.Wait(ctx); err != nil {
		t.Fatalf("domain completion gate never opened: %v", err)
	}
	ledger := store.NewMobilizeCompletionStore(3)
	ledger.Acknowledge("north")
	short2, cancelShort2 := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelShort2()
	if ledger.Wait(short2) == nil {
		t.Fatal("mobilize ledger completed after one acknowledgement")
	}
	ledger.Acknowledge("south")
	ledger.Acknowledge("subsea")
	if err := ledger.Wait(ctx); err != nil {
		t.Fatalf("mobilize ledger never completed: %v", err)
	}
	want := []string{"north", "south", "subsea"}
	got := worker.CoordinateSpans(ctx, want)
	sort.Strings(got)
	if len(got) != 3 {
		t.Fatalf("worker coordinator returned %v", got)
	}
	session := service.RunMobilizeSession(ctx, want)
	sort.Strings(session)
	if len(session) != 3 {
		t.Fatalf("service session returned %v", session)
	}
}
