package store

import (
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"
	"reflect"
)

type SafetyRegistry struct {
	providers map[string]domain.SafetyProvider
}

func NewSafetyRegistry() *SafetyRegistry {
	return &SafetyRegistry{providers: map[string]domain.SafetyProvider{}}
}
func providerPresent(p domain.SafetyProvider) bool {
	if p == nil {
		return false
	}
	value := reflect.ValueOf(p)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return !value.IsNil()
	default:
		return true
	}
}
func (r *SafetyRegistry) Register(name string, p domain.SafetyProvider) bool {
	if !providerPresent(p) {
		return false
	}
	r.providers[name] = p
	return true
}
func (r *SafetyRegistry) Has(name string) bool { return providerPresent(r.providers[name]) }
