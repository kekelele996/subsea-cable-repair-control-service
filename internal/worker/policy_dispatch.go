package worker

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

type SafetyJob struct {
	Plan     domain.RepairPlan
	Provider domain.SafetyProvider
}

func AcceptSafetyJob(job SafetyJob) error {
	if job.Provider == nil {
		return domain.ErrSafetyUnavailable
	}
	return nil
}
