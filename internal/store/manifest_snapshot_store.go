package store

import (
	"sync"

	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
)

type ManifestSnapshotStore struct {
	mu    sync.Mutex
	items map[string]domain.ManifestSnapshot
}

func NewManifestSnapshotStore() *ManifestSnapshotStore {
	return &ManifestSnapshotStore{items: map[string]domain.ManifestSnapshot{}}
}

func copyStoredManifest(in domain.ManifestSnapshot) domain.ManifestSnapshot {
	out := in
	out.Spans = append([]string(nil), in.Spans...)
	out.Contacts = append([]string(nil), in.Contacts...)
	out.RequiredChecks = make(map[string][]string, len(in.RequiredChecks))
	for span, checks := range in.RequiredChecks {
		out.RequiredChecks[span] = append([]string(nil), checks...)
	}
	return out
}

func (s *ManifestSnapshotStore) Save(snapshot domain.ManifestSnapshot) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[snapshot.PlanID] = copyStoredManifest(snapshot)
}

func (s *ManifestSnapshotStore) Load(planID string) domain.ManifestSnapshot {
	s.mu.Lock()
	defer s.mu.Unlock()
	return copyStoredManifest(s.items[planID])
}
