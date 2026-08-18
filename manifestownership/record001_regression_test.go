package manifestcheck

import (
	"reflect"
	"sync"
	"testing"

	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/service"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/store"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/worker"
)

func manifestFixture() domain.ManifestSnapshot {
	spans := make([]string, 2, 4)
	copy(spans, []string{"north", "south"})
	contacts := make([]string, 1, 3)
	contacts[0] = "control-room"
	checks := map[string][]string{"north": {"pressure", "corrosion"}}
	return domain.ManifestSnapshot{PlanID: "rig-17", Spans: spans, Contacts: contacts, RequiredChecks: checks}
}

func TestManifestSnapshotHasIndependentMemory(t *testing.T) {
	source := manifestFixture()
	domainView := domain.CopyManifestSnapshot(source)
	source.Spans[0] = "mutated-input"
	source.RequiredChecks["north"][0] = "mutated-check"
	if domainView.Spans[0] != "north" || domainView.RequiredChecks["north"][0] != "pressure" {
		t.Fatalf("domain snapshot changed with caller input: %+v", domainView)
	}

	storedInput := manifestFixture()
	manifests := store.NewManifestSnapshotStore()
	manifests.Save(storedInput)
	storedInput.Contacts[0] = "mutated-contact"
	firstRead := manifests.Load("rig-17")
	firstRead.Spans[0] = "mutated-read"
	secondRead := manifests.Load("rig-17")
	if secondRead.Contacts[0] != "control-room" || secondRead.Spans[0] != "north" {
		t.Fatalf("store snapshot ownership was lost: %+v", secondRead)
	}

	serviceInput := manifestFixture()
	retained := domain.CopyManifestSnapshot(serviceInput)
	emergency := service.EmergencyManifestView(retained, "backup", "marine-desk")
	if !reflect.DeepEqual(retained.Spans, []string{"north", "south"}) || len(retained.RequiredChecks["backup"]) != 0 {
		t.Fatalf("retained service view changed after emergency expansion: %+v", retained)
	}
	if len(emergency.Spans) != 3 || len(emergency.RequiredChecks["backup"]) != 1 {
		t.Fatalf("emergency view incomplete: %+v", emergency)
	}

	mobilizeInput := manifestFixture()
	delivered := <-worker.MobilizeManifest(mobilizeInput)
	start := make(chan struct{})
	var participants sync.WaitGroup
	participants.Add(2)
	go func() {
		defer participants.Done()
		<-start
		for i := 0; i < 1000; i++ {
			delivered.Spans[0] = "consumer-overwrite"
			delivered.RequiredChecks["north"][0] = "consumer-overwrite"
		}
	}()
	go func() {
		defer participants.Done()
		<-start
		for i := 0; i < 1000; i++ {
			_ = mobilizeInput.Spans[0]
			_ = mobilizeInput.RequiredChecks["north"][0]
		}
	}()
	close(start)
	participants.Wait()
	if mobilizeInput.Spans[0] != "north" || mobilizeInput.RequiredChecks["north"][0] != "pressure" {
		t.Fatalf("worker delivery exposed producer memory: %+v", mobilizeInput)
	}
}
