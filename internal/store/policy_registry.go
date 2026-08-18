package store

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

type SafetyRegistry struct {
	providers map[string]domain.SafetyProvider
}

func NewSafetyRegistry() *SafetyRegistry {
	return &SafetyRegistry{providers: map[string]domain.SafetyProvider{}}
}
func (r *SafetyRegistry) Register(name string, p domain.SafetyProvider) bool {
	if p == nil {
		return false
	}
	r.providers[name] = p
	return true
}
func (r *SafetyRegistry) Has(name string) bool { return r.providers[name] != nil }
