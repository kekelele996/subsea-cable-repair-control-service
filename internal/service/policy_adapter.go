package service

import (
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"reflect"
)

func EvaluateProvider(plan domain.RepairPlan, p domain.SafetyProvider) ([]domain.Hazard, error) {
	if p == nil {
		return nil, domain.ErrSafetyUnavailable
	}
	value := reflect.ValueOf(p)
	if value.Kind() == reflect.Pointer && value.IsNil() {
		return nil, domain.ErrSafetyUnavailable
	}
	return p.Evaluate(plan)
}
