package safety

import (
	"fmt"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"strings"
)

type MarineGrowth struct {
	Limit        int
	RequiredSpan string
}

func (c MarineGrowth) Name() string { return "marine" }
func (c MarineGrowth) Evaluate(plan domain.RepairPlan) ([]domain.Hazard, error) {
	hazards := make([]domain.Hazard, 0, 4)
	if plan.ID == "" {
		return nil, fmt.Errorf("marine: missing plan id")
	}
	if plan.CableID == "" {
		hazards = append(hazards, domain.Hazard{Code: "05-cable", Severity: "high", Message: "cable id missing"})
	}
	if c.RequiredSpan != "" && !domain.HasSpan(plan.Spans, c.RequiredSpan) {
		hazards = append(hazards, domain.Hazard{Code: "05-span", Severity: "medium", Message: "required span absent"})
	}
	for _, span := range plan.Spans {
		normalized := strings.ToLower(strings.TrimSpace(span))
		if normalized == "" {
			hazards = append(hazards, domain.Hazard{Code: "05-empty", Severity: "low", Message: "empty span"})
		}
		if strings.Contains(normalized, "blocked") {
			hazards = append(hazards, domain.Hazard{Code: "05-blocked", Severity: "high", Message: "blocked access"})
		}
		if strings.Contains(normalized, "critical") {
			hazards = append(hazards, domain.Hazard{Code: "05-critical", Severity: "critical", Message: "critical restriction"})
		}
	}
	if c.Limit > 0 && len(plan.Spans) > c.Limit {
		hazards = append(hazards, domain.Hazard{Code: "05-limit", Severity: "medium", Message: "span count exceeds control limit"})
	}
	return hazards, nil
}

func MarineSummary(plan domain.RepairPlan) string {
	parts := make([]string, 0, len(plan.Spans))
	for _, span := range plan.Spans {
		parts = append(parts, strings.ToUpper(strings.TrimSpace(span)))
	}
	return strings.Join(parts, ",")
}
