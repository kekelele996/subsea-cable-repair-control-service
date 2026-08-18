package store

import (
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"sync"
)

type ManifestStore struct {
	mu    sync.Mutex
	spans map[string][]string
}

func NewManifestStore() *ManifestStore { return &ManifestStore{spans: map[string][]string{}} }
func (s *ManifestStore) Save(plan string, spans []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.spans[plan] = domain.CloneStrings(spans)
}
func (s *ManifestStore) Read(plan string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return domain.CloneStrings(s.spans[plan])
}
