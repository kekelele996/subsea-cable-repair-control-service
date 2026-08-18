package domain

// FinalizationOutcome combines the primary finalization result with a cleanup
// result. The primary failure is the authoritative outcome and must never be
// masked: if finalization failed, callers (and the API) must observe that
// failure regardless of what cleanup reports. A cleanup error surfaces only when
// finalization itself succeeded, so a failed write is never turned into a
// success by a later cleanup step.
func FinalizationOutcome(primary, cleanup error) error {
	if primary != nil {
		return primary
	}
	return cleanup
}
