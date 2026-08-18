package store

import (
	"context"
	"sync"
)

type MobilizeCompletionStore struct {
	mu       sync.Mutex
	expected int
	acked    map[string]bool
	done     chan struct{}
	once     sync.Once
}

func NewMobilizeCompletionStore(expected int) *MobilizeCompletionStore {
	return &MobilizeCompletionStore{expected: expected, acked: map[string]bool{}, done: make(chan struct{})}
}
func (s *MobilizeCompletionStore) Acknowledge(span string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.acked[span] = true
	if len(s.acked) == s.expected {
		s.once.Do(func() { close(s.done) })
	}
}
func (s *MobilizeCompletionStore) Wait(ctx context.Context) error {
	select {
	case <-s.done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
