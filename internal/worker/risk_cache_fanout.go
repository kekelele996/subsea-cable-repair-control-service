package worker

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func FanoutRisk(hazards []domain.Hazard) ([]domain.Hazard, []domain.Hazard) {
	return domain.CloneHazards(hazards), domain.CloneHazards(hazards)
}
