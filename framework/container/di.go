package container

import "github.com/samber/do"

var injector = do.New()

type Provider[T any] func() (T, error)

func Provide(provider Provider[any]) {
	do.Provide(injector, wrapProvider(provider))
}

func ProvideNamed(name string, provider Provider[any]) {
	do.ProvideNamed(injector, name, wrapProvider(provider))
}

func Invoke[T any]() (T, error) {
	return do.Invoke[T](injector)
}

func MustInvoke[T any]() T {
	return do.MustInvoke[T](injector)
}

func InvokeNamed[T any](name string) (T, error) {
	return do.InvokeNamed[T](injector, name)
}

func MustInvokeNamed[T any](name string) T {
	return do.MustInvokeNamed[T](injector, name)
}

func wrapProvider(provider Provider[any]) func(*do.Injector) (any, error) {
	return func(*do.Injector) (any, error) {
		return provider()
	}
}
