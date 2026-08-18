package store

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

type RiskCache struct {
	hazards []domain.Hazard
}

func (c *RiskCache) Put(hazards []domain.Hazard) {
	c.hazards = append(c.hazards[:0], hazards...)
}

func (c *RiskCache) Read() []domain.Hazard {
	snapshot := make([]domain.Hazard, len(c.hazards))
	copy(snapshot, c.hazards)
	return snapshot
}
