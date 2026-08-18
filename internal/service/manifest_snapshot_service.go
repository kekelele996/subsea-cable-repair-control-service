package service

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func EmergencyManifestView(current domain.ManifestSnapshot, span, contact string) domain.ManifestSnapshot {
	next := current
	next.Spans = append(append([]string(nil), current.Spans...), span)
	next.Contacts = append(append([]string(nil), current.Contacts...), contact)
	next.RequiredChecks = make(map[string][]string, len(current.RequiredChecks)+1)
	for name, checks := range current.RequiredChecks {
		next.RequiredChecks[name] = append([]string(nil), checks...)
	}
	next.RequiredChecks[span] = []string{"gas-test"}
	return next
}
