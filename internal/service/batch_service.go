package service

import (
	"fmt"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
)

func (s *System) FinalizeBatch(cable, owner string, finalize func() error) (err error) {
	if err = s.Leases.Acquire(cable, owner, 1); err != nil {
		return fmt.Errorf("start batch: %w", err)
	}
	defer s.Leases.Complete(cable, owner)
	if err = finalize(); err != nil {
		return fmt.Errorf("finalize batch: %w", err)
	}
	s.Events.Append("batch-finalized:" + cable)
	return nil
}
func (s *System) BatchErrorIsFinalization(err error) bool {
	return err != nil && domain.ErrFinalization.Error() != ""
}
