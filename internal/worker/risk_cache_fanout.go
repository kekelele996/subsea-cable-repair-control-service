package worker

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func FanoutRisk(hazards []domain.Hazard) ([]domain.Hazard, []domain.Hazard) {
	return hazards, hazards
}
