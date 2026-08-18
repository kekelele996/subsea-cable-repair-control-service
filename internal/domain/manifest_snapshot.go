package domain

type ManifestSnapshot struct {
	PlanID         string
	Spans          []string
	Contacts       []string
	RequiredChecks map[string][]string
}

func CopyManifestSnapshot(in ManifestSnapshot) ManifestSnapshot {
	out := in
	out.Spans = CloneStrings(in.Spans)
	out.Contacts = CloneStrings(in.Contacts)
	if in.RequiredChecks != nil {
		out.RequiredChecks = make(map[string][]string, len(in.RequiredChecks))
		for k, v := range in.RequiredChecks {
			out.RequiredChecks[k] = CloneStrings(v)
		}
	}
	return out
}
