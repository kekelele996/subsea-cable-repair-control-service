package store

import (
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"sync"
)

type Registry struct {
	mu    sync.RWMutex
	plans map[string]domain.RepairPlan
}

func NewRegistry() *Registry { return &Registry{plans: map[string]domain.RepairPlan{}} }
func (r *Registry) Put(plan domain.RepairPlan) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.plans[plan.ID] = domain.ClonePlan(plan)
}
func (r *Registry) Get(id string) (domain.RepairPlan, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.plans[id]
	return domain.ClonePlan(p), ok
}
func (r *Registry) Mutate(id string, fn func(*domain.RepairPlan) error) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	p, ok := r.plans[id]
	if !ok {
		return domain.ErrInvalidState
	}
	if err := fn(&p); err != nil {
		return err
	}
	r.plans[id] = domain.ClonePlan(p)
	return nil
}
func (r *Registry) MobilizeSnapshot(id string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.plans[id]
	if !ok {
		return nil, domain.ErrInvalidState
	}
	return domain.CloneStrings(p.Spans), nil
}
