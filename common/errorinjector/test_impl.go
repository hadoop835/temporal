//go:build errorinjector

package errorinjector

import (
	"sync"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(func() ErrorInjector { return newTestErrorInjector() }),
)

type (
	ErrorInjector interface {
		// private accessors; access must go through package-level Get/Set
		get(string) (any, bool)
		set(string, any)
	}

	errorInjectorImpl struct {
		m sync.Map
	}
)

func Get[T any](ei ErrorInjector, key string) (T, bool) {
	if val, ok := ei.get(key); ok {
		// this is only used in test so we want to panic on type mismatch:
		return val.(T), ok
	}
	var zero T
	return zero, false
}

func Set[T any](ei ErrorInjector, key string, val T) {
	ei.set(key, val)
}

func newTestErrorInjector() *errorInjectorImpl {
	return &errorInjectorImpl{}
}

func (ei *errorInjectorImpl) get(key string) (any, bool) {
	val, ok := ei.m.Load(key)
	return val, ok
}

func (ei *errorInjectorImpl) set(key string, val any) {
	ei.m.Store(key, val)
}
