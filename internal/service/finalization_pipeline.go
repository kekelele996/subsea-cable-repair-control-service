package service

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

// FinalizeRepair runs finalization and then a cleanup step, returning the
// authoritative outcome. The finalization error is never overwritten by
// cleanup: a failed write to the store stays a failure at the API even when
// cleanup (slot release, event append) succeeds. A cleanup error surfaces only
// when finalization itself succeeded, so cleanup failures are still reported
// instead of being swallowed.
func FinalizeRepair(finalize, cleanup func() error) (err error) {
	primary := finalize()
	cleanupErr := cleanup()
	err = domain.FinalizationOutcome(primary, cleanupErr)
	return
}
