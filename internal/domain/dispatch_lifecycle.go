package domain

import (
	"context"
	"sync"
)

type CompletionGate struct {
	mu        sync.Mutex
	remaining int
	done      chan struct{}
}

func NewCompletionGate(expected int) *CompletionGate {
	return &CompletionGate{remaining: expected, done: make(chan struct{})}
}
func (g *CompletionGate) Complete() {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.remaining == 0 {
		return
	}
	g.remaining--
	if g.remaining == 0 {
		close(g.done)
	}
}
func (g *CompletionGate) Wait(ctx context.Context) error {
	g.mu.Lock()
	done := g.done
	g.mu.Unlock()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
