package worker

import (
	"context"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/service"
)

func Mobilize(ctx context.Context, s *service.System, planID string) ([]string, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	spans, err := s.Registry.MobilizeSnapshot(planID)
	if err != nil {
		return nil, err
	}
	jobs := make([]domain.RepairJob, 0, len(spans))
	for _, span := range spans {
		jobs = append(jobs, domain.RepairJob{ID: planID + ":" + span, PlanID: planID, Spans: []string{span}})
	}
	return Coordinate(ctx, jobs)
}
