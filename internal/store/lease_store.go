package store

import (
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"sync"
	"time"
)

type LeaseStore struct {
	mu     sync.Mutex
	leases map[string]domain.Lease
	now    func() time.Time
}

func NewLeaseStore(now func() time.Time) *LeaseStore {
	return &LeaseStore{leases: map[string]domain.Lease{}, now: now}
}
func (s *LeaseStore) Acquire(cable, owner string, ttl time.Duration) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if old, ok := s.leases[cable]; ok && old.ExpiresAt.After(s.now()) && old.Owner != owner {
		return domain.ErrLeaseConflict
	}
	s.leases[cable] = domain.Lease{CableID: cable, Owner: owner, ExpiresAt: s.now().Add(ttl)}
	return nil
}
func (s *LeaseStore) Complete(cable, owner string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if old, ok := s.leases[cable]; ok && old.Owner == owner {
		delete(s.leases, cable)
	}
}
func (s *LeaseStore) Held(cable string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	old, ok := s.leases[cable]
	return ok && old.ExpiresAt.After(s.now())
}
