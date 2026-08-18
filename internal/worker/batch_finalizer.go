package worker

type BatchFinalizer struct{ Acknowledged bool }

// Execute runs a batch and acknowledges it only on success. A batch whose
// finalization failed must not be acknowledged: acking it would let downstream
// consumers treat a failed batch as committed and mark its task done. The error
// is returned unchanged so the caller observes the real failure.
func (f *BatchFinalizer) Execute(run func() error) (err error) {
	err = run()
	if err == nil {
		f.Acknowledged = true
	}
	return
}
