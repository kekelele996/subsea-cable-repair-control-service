package store

import (
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"sync"
)

type RiskCache struct {
	mu      sync.RWMutex
	hazards map[string][]domain.Hazard
}

func NewRiskCache() *RiskCache { return &RiskCache{hazards: map[string][]domain.Hazard{}} }
func (c *RiskCache) Put(plan string, hazards []domain.Hazard) {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]domain.Hazard, len(hazards))
	copy(out, hazards)
	c.hazards[plan] = out
}
func (c *RiskCache) Get(plan string) []domain.Hazard {
	c.mu.RLock()
	defer c.mu.RUnlock()
	values := c.hazards[plan]
	out := make([]domain.Hazard, len(values))
	copy(out, values)
	return out
}
