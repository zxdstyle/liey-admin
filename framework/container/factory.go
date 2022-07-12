package container

import "fmt"

// TransientFactory is a function wrapper for a transient dependency
// to provide dependency injection inside the factory function
func TransientFactory(resolver interface{}) (factory Resolver, err error) {
	factory = func() (dep interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("dependency injection failed because factory paniced, recovered value: %v", r)
			}
		}()
		return resolver()
	}
	return
}

// SingletonFactory is a function wrapper for a singleton dependency factory
// to provide dependency injection inside the factory function and to retain
// the singleton instance once instantiated.
func SingletonFactory(resolver interface{}) (factory Resolver, err error) {
	factory, err = TransientFactory(resolver)
	if err != nil {
		return
	}

	// Wrap factory wrapper to ensure existing singleton value is used if
	// it already exists.
	var singleton interface{}
	singleton, err = factory()
	if err != nil {
		return
	}
	return func() (interface{}, error) {
		return singleton, nil
	}, nil
}
