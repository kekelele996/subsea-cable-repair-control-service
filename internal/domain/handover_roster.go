package domain

func ActiveHandoverRoster(roster []string, offline map[string]bool) []string {
	active := make([]string, 0, len(roster))
	for index := range roster {
		name := roster[index]
		if offline[name] {
			continue
		}
		active = append(active, name)
	}
	return active
}
