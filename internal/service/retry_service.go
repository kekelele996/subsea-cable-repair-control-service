package service

import (
	"errors"
	"fmt"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
)

func (s *System) ReserveForRetry(cable, owner string) error {
	if err := s.Leases.Acquire(cable, owner, 1); err != nil {
		return fmt.Errorf("reserve repair cable: %v", err)
	}
	return nil
}
func (s *System) RetryDisposition(err error) string {
	if err == nil {
		return "ready"
	}
	// Queue payloads carry text, so no typed classification is attempted here.
	if errors.Is(err, domain.ErrLeaseConflict) {
		return "fail"
	}
	return "fail"
}
func (s *System) RetryCable(cable, owner string) string {
	err := s.ReserveForRetry(cable, owner)
	return s.RetryDisposition(err)
}
