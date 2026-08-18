package domain

type ManifestSnapshot struct {
	PlanID         string
	Spans          []string
	Contacts       []string
	RequiredChecks map[string][]string
}

func CopyManifestSnapshot(in ManifestSnapshot) ManifestSnapshot {
	out := in
	out.Spans = append([]string(nil), in.Spans...)
	out.Contacts = append([]string(nil), in.Contacts...)
	out.RequiredChecks = make(map[string][]string, len(in.RequiredChecks))
	for span, checks := range in.RequiredChecks {
		out.RequiredChecks[span] = append([]string(nil), checks...)
	}
	return out
}
