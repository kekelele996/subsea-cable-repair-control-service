package worker

import (
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"reflect"
)

type SafetyJob struct {
	Plan     domain.RepairPlan
	Provider domain.SafetyProvider
}

func AcceptSafetyJob(job SafetyJob) error {
	if job.Provider == nil {
		return domain.ErrSafetyUnavailable
	}
	value := reflect.ValueOf(job.Provider)
	if value.Kind() == reflect.Pointer && value.IsNil() {
		return domain.ErrSafetyUnavailable
	}
	return nil
}
