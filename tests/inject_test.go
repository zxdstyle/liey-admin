package tests

import (
	"fmt"
	"github.com/samber/do"
	"testing"
)

var injector = do.New()

func TestName(t *testing.T) {
	fmt.Printf("%T", injector)
	do.Provide(injector, NewCarService)
	do.Provide(injector, NewEngineService)

	car1 := do.MustInvoke[*CarService](injector)
	car := do.MustInvoke[*CarService](injector)

	car1.Start()
	car.Start()
	injector.ListProvidedServices()
}

type EngineService interface{}

func NewEngineService(i *do.Injector) (EngineService, error) {
	return &engineServiceImplem{}, nil
}

type engineServiceImplem struct{}

func (c *engineServiceImplem) HealthCheck() error {
	return fmt.Errorf("engine broken")
}

func NewCarService(i *do.Injector) (*CarService, error) {
	engine := do.MustInvoke[EngineService](i)
	car := CarService{Engine: engine}
	fmt.Println("test")
	return &car, nil
}

type CarService struct {
	Engine EngineService
}

func (c *CarService) Start() {
	println("car starting")
}

func (c *CarService) Shutdown() error {
	println("car stopped")
	return nil
}
