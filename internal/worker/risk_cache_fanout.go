package worker

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func FanoutRisk(hazards []domain.Hazard) ([]domain.Hazard, []domain.Hazard) {
	left := make([]domain.Hazard, len(hazards))
	right := make([]domain.Hazard, len(hazards))
	copy(left, hazards)
	copy(right, hazards)
	return left, right
}
