package service

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func EmergencyManifestView(current domain.ManifestSnapshot, span, contact string) domain.ManifestSnapshot {
	current.Spans = append(current.Spans, span)
	current.Contacts = append(current.Contacts, contact)
	current.RequiredChecks[span] = append(current.RequiredChecks[span], "gas-test")
	return current
}
