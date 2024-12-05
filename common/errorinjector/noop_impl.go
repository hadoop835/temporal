//go:build !errorinjector

package errorinjector

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(func() (ei ErrorInjector) { return }),
)

type (
	ErrorInjector struct{}
)

func Get[T any](ei ErrorInjector, key string) (T, bool) {
	var zero T
	return zero, false
}
