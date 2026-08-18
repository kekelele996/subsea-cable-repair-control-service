package store

import (
	"context"
	"sync"
)

type MobilizeCompletionStore struct {
	mu       sync.Mutex
	expected int
	acked    map[string]bool
	done      chan struct{}
}

func NewMobilizeCompletionStore(expected int) *MobilizeCompletionStore {
	return &MobilizeCompletionStore{expected: expected, acked: map[string]bool{}, done: make(chan struct{})}
}
func (s *MobilizeCompletionStore) Acknowledge(span string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.done == nil || s.acked[span] {
		return
	}
	s.acked[span] = true
	if len(s.acked) >= s.expected {
		close(s.done)
		s.done = nil
	}
}
func (s *MobilizeCompletionStore) Wait(ctx context.Context) error {
	s.mu.Lock()
	done := s.done
	s.mu.Unlock()
	if done != nil {
		select {
		case <-done:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}
