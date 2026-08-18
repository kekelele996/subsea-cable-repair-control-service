package service_test

import (
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/service"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/store"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/worker"
	"testing"
)

func TestRiskCacheOperationalNoteDoesNotMutateStoredHazards(t *testing.T) {
	s := service.NewRiskCacheService(store.NewRiskCache())
	s.Record("p10", []domain.Hazard{{Code: "barrier", Severity: "low"}})
	_ = worker.AttachOperationalNote(s, "p10")
	if got := s.Current("p10"); len(got) != 1 || got[0].Code != "barrier" {
		t.Fatalf("cache polluted: %#v", got)
	}
}
