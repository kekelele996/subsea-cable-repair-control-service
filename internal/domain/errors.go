package domain

import "errors"

var (
	ErrLeaseConflict     = errors.New("cable lease conflict")
	ErrSafetyUnavailable = errors.New("safety provider unavailable")
	ErrUnsafeComplete    = errors.New("unsafe complete")
	ErrCancelled         = errors.New("repair was cancelled")
	ErrFinalization      = errors.New("batch finalization failed")
	ErrInvalidState      = errors.New("invalid plan state")
)
