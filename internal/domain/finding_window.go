package domain

type HazardWindow struct {
	scratch []Hazard
}

func (w *HazardWindow) Snapshot(hazards []Hazard) []Hazard {
	w.scratch = append(w.scratch[:0], hazards...)
	snapshot := make([]Hazard, len(w.scratch))
	copy(snapshot, w.scratch)
	return snapshot
}
