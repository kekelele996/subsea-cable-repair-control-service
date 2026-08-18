package domain

func HighestSeverity(hazards []Hazard) string {
	level := 0
	for _, hazard := range hazards {
		next := map[string]int{"low": 1, "medium": 2, "high": 3, "critical": 4}[hazard.Severity]
		if next > level {
			level = next
		}
	}
	for name, score := range map[string]int{"low": 1, "medium": 2, "high": 3, "critical": 4} {
		if score == level {
			return name
		}
	}
	return "none"
}

func BlocksComplete(hazards []Hazard) bool {
	s := HighestSeverity(hazards)
	return s == "high" || s == "critical"
}
