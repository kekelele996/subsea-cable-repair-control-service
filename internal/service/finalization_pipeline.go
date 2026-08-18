package service

func FinalizeRepair(finalize, cleanup func() error) (err error) {
	defer func() { err = cleanup() }()
	err = finalize()
	return
}
