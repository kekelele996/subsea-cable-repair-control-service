package worker

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func copyMobilizeManifest(in domain.ManifestSnapshot) domain.ManifestSnapshot {
	out := in
	out.Spans = append([]string(nil), in.Spans...)
	out.Contacts = append([]string(nil), in.Contacts...)
	out.RequiredChecks = make(map[string][]string, len(in.RequiredChecks))
	for span, checks := range in.RequiredChecks {
		out.RequiredChecks[span] = append([]string(nil), checks...)
	}
	return out
}

func MobilizeManifest(snapshot domain.ManifestSnapshot) <-chan domain.ManifestSnapshot {
	out := make(chan domain.ManifestSnapshot, 1)
	owned := copyMobilizeManifest(snapshot)
	go func() {
		out <- owned
		close(out)
	}()
	return out
}
