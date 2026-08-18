package service

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func EvaluateProvider(plan domain.RepairPlan, p domain.SafetyProvider) ([]domain.Hazard, error) {
	if p == nil {
		return nil, domain.ErrSafetyUnavailable
	}
	return p.Evaluate(plan)
}
