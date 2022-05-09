package publisher

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
)

type Config struct {
	source []byte
}

func NewConfigPublisher(source []byte) Config {
	return Config{source: source}
}

func (c Config) Source() []byte {
	return c.source
}

func (c Config) Publish(name string) error {
	path := fmt.Sprintf("config/%s.yaml", name)
	if gfile.Exists(path) {
		return fmt.Errorf("%s already exists", path)
	}

	return gfile.PutBytes(path, c.Source())
}
