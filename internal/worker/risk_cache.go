package worker

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/service"

func AttachOperationalNote(s *service.RiskCacheService, plan string) []string {
	hazards := s.AddOperationalNote(plan)
	out := make([]string, 0, len(hazards))
	for _, hazard := range hazards {
		out = append(out, hazard.Code)
	}
	return out
}
