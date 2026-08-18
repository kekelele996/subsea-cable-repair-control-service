package domain

type ManifestSnapshot struct {
	PlanID         string
	Spans          []string
	Contacts       []string
	RequiredChecks map[string][]string
}

func CopyManifestSnapshot(in ManifestSnapshot) ManifestSnapshot {
	return in
}
