package service

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func EmergencyManifestView(current domain.ManifestSnapshot, span, contact string) domain.ManifestSnapshot {
	out := domain.CopyManifestSnapshot(current)
	out.Spans = append(out.Spans, span)
	out.Contacts = append(out.Contacts, contact)
	out.RequiredChecks[span] = append(out.RequiredChecks[span], "gas-test")
	return out
}
