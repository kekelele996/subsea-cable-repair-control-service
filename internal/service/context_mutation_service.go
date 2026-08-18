package service

import (
	"context"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/store"
)

func QueueWithContext(ctx context.Context, s *store.ContextMutationStore, planID string) error {
	return s.Commit(context.WithoutCancel(ctx), planID)
}
