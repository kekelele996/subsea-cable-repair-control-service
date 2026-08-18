package domain

type HazardWindow struct{ scratch []Hazard }

func (w *HazardWindow) Snapshot(hazards []Hazard) []Hazard {
	w.scratch = append(w.scratch[:0], hazards...)
	return CloneHazards(w.scratch)
}
