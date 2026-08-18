package service

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func (s *System) StageManifest(id string) error {
	spans, err := s.Registry.MobilizeSnapshot(id)
	if err != nil {
		return err
	}
	s.Manifests.Save(id, spans)
	return nil
}
func (s *System) AddEmergencySpan(id, span string) []string {
	spans := s.Manifests.Read(id)
	spans = append(spans, span)
	return domain.CloneStrings(spans)
}
func (s *System) StoredManifest(id string) []string { return s.Manifests.Read(id) }
