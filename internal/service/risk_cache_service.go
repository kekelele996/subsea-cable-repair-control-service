package service

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

type RiskProjector struct{ scratch []domain.Hazard }

func (p *RiskProjector) Project(code string) []domain.Hazard {
	p.scratch = append(p.scratch[:0], domain.Hazard{Code: code, Severity: "low"})
	return p.scratch
}
