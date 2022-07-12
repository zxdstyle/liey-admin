package container

import (
	"fmt"
	"github.com/zxdstyle/structures"
	"reflect"
)

var instances = structures.NewMap[string, func() (interface{}, error)]()

func Singleton[V any](resolver func() (V, error)) error {
	factory, err := SingletonFactory(resolver)
	if err != nil {
		return err
	}
	depFactoryType := reflect.TypeOf(resolver)
	key := reflectTypeKey(depFactoryType.Out(0))
	instances.Set(key, factory)
	return nil
}

func Register(resolver Resolver) error {
	factory, err := TransientFactory(resolver)
	if err != nil {
		return err
	}
	depFactoryType := reflect.TypeOf(resolver)
	instances.Set(reflectTypeKey(depFactoryType.Out(0)), factory)
	return nil
}

// InjectT tries to resolve a dependency by its type and optionally its name
// If the dependency is unknown, ErrorDepFactoryNotFound is returned
func InjectT(depType reflect.Type, name ...string) (interface{}, error) {
	factory := instances.Get(reflectTypeKey(depType))
	if factory == nil {
		return nil, ErrorDepFactoryNotFound
	}
	return factory()
}

func Inject[V any](out interface{}) error {
	argType := reflect.TypeOf(out)
	if argType.Kind() != reflect.Pointer {
		return ErrorDepNotAPointer
	}
	factory := instances.Get(reflectTypeKey(argType.Elem()))
	if factory == nil {
		return ErrorDepFactoryNotFound
	}
	dep, err := factory()
	if err != nil {
		return err
	}
	reflect.ValueOf(out).Elem().Set(reflect.ValueOf(dep))
	return nil
}

// Call will attempt to resolve all arguments of the function and then call it
func Call(fn interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("dependency injection failed because factory paniced, recovered value: %v", r)
		}
	}()
	args, err := injectFuncArgs(fn)
	if err != nil {
		return
	}
	resVal := reflect.ValueOf(fn).Call(args)
	if len(resVal) > 0 {
		if err, ok := resVal[0].Interface().(error); ok && err != nil {
			return err
		}
	}
	return
}
