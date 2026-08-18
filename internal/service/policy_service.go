package service

import (
	"fmt"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"reflect"
)

func providerUnavailable(provider domain.SafetyProvider) bool {
	if provider == nil {
		return true
	}
	v := reflect.ValueOf(provider)
	return v.Kind() == reflect.Ptr && v.IsNil()
}
func (s *System) EvaluateComplete(id string, provider domain.SafetyProvider) (domain.CompletionDecision, error) {
	p, ok := s.Registry.Get(id)
	if !ok {
		return domain.CompletionDecision{}, domain.ErrInvalidState
	}
	if providerUnavailable(provider) {
		return domain.CompletionDecision{PlanID: id, Allowed: false, Reason: "safety unavailable"}, domain.ErrSafetyUnavailable
	}
	hazards, err := provider.Evaluate(p)
	if err != nil {
		return domain.CompletionDecision{}, fmt.Errorf("evaluate complete: %w", err)
	}
	if domain.BlocksComplete(hazards) {
		return domain.CompletionDecision{PlanID: id, Allowed: false, Reason: "blocking hazard"}, nil
	}
	return domain.CompletionDecision{PlanID: id, Allowed: true}, nil
}
