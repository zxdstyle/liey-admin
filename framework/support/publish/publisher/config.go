package publisher

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
)

type Config struct {
	sources []Publishable
}

func NewConfigPublisher(sources []Publishable) Config {
	return Config{sources: sources}
}

func (c Config) Sources() []Publishable {
	return c.sources
}

func (c Config) Publish() error {
	for _, source := range c.sources {
		path := source.Path
		if gfile.Exists(path) {
			return fmt.Errorf("%s already exists", path)
		}

		if err := gfile.PutBytes(path, source.Content); err != nil {
			return err
		}
	}
	return nil
}
