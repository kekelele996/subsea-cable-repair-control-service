package service

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

type RiskCacheService struct {
	cache interface {
		Put(string, []domain.Hazard)
		Get(string) []domain.Hazard
	}
}

func NewRiskCacheService(cache interface {
	Put(string, []domain.Hazard)
	Get(string) []domain.Hazard
}) *RiskCacheService {
	return &RiskCacheService{cache: cache}
}
func (s *RiskCacheService) Record(plan string, hazards []domain.Hazard) {
	s.cache.Put(plan, hazards)
}
func (s *RiskCacheService) AddOperationalNote(plan string) []domain.Hazard {
	hazards := s.cache.Get(plan)
	hazards = append(hazards, domain.Hazard{Code: "ops-note", Severity: "low", Message: "operator note"})
	return hazards
}
func (s *RiskCacheService) Current(plan string) []domain.Hazard { return s.cache.Get(plan) }
