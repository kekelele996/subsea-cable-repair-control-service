package store

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

type CompleteRepository struct{ Plan domain.RepairPlan }

func NewCompleteRepository(plan domain.RepairPlan) *CompleteRepository {
	return &CompleteRepository{Plan: plan}
}
func (r *CompleteRepository) Commit(expectedRevision int) error {
	r.Plan.State = domain.RepairCompleted
	r.Plan.Revision++
	return nil
}
