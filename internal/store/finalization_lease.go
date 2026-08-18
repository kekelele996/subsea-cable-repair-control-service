package store

type FinalizationLease struct{ held bool }

func NewFinalizationLease() *FinalizationLease { return &FinalizationLease{held: true} }
func (l *FinalizationLease) Run(finalize func() error) (err error) {
	// The lease guards a repair-vessel slot for the duration of finalization.
	// Once finalization has run to completion it must be released whether it
	// succeeded or failed: a failed batch should not strand the slot. The
	// finalization error is preserved verbatim so cleanup never masks failure.
	defer func() { l.held = false }()
	return finalize()
}
func (l *FinalizationLease) Held() bool { return l.held }
