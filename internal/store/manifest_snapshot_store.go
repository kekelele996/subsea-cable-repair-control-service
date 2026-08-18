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

func (s *ManifestSnapshotStore) Save(snapshot domain.ManifestSnapshot) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[snapshot.PlanID] = domain.CopyManifestSnapshot(snapshot)
}

func (s *ManifestSnapshotStore) Load(planID string) domain.ManifestSnapshot {
	s.mu.Lock()
	defer s.mu.Unlock()
	return domain.CopyManifestSnapshot(s.items[planID])
}
