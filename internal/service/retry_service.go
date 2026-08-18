package service

import (
	"errors"
	"fmt"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
)

func (s *System) ReserveForRetry(cable, owner string) error {
	if err := s.Leases.Acquire(cable, owner, 1); err != nil {
		return fmt.Errorf("reserve repair cable: %w", err)
	}
	return nil
}
func (s *System) RetryDisposition(err error) string {
	if errors.Is(err, domain.ErrLeaseConflict) {
		return "defer"
	}
	return "fail"
}
func (s *System) RetryCable(cable, owner string) string {
	err := s.ReserveForRetry(cable, owner)
	return s.RetryDisposition(err)
}
