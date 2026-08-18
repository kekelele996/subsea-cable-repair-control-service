package domain

func CloneStrings(values []string) []string {
	if values == nil {
		return nil
	}
	out := make([]string, len(values))
	copy(out, values)
	return out
}

func CloneHazards(values []Hazard) []Hazard {
	if values == nil {
		return nil
	}
	out := make([]Hazard, len(values))
	copy(out, values)
	return out
}

func ClonePlan(in RepairPlan) RepairPlan {
	in.Spans = CloneStrings(in.Spans)
	return in
}

func HasSpan(spans []string, wanted string) bool {
	for _, span := range spans {
		if span == wanted {
			return true
		}
	}
	return false
}
