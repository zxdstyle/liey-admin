package serializer

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
)

type UintSerializer struct {
}

func (s UintSerializer) Serialize(payload interface{}) ([]byte, error) {
	return gvar.New(payload).Bytes(), nil
}

func (s UintSerializer) UnSerialize(payload []byte, pointer interface{}) error {
	val, ok := pointer.(*uint)
	if !ok {
		return fmt.Errorf("UintSerializer receiver must be `*uint`")
	}

	*val = gvar.New(payload).Uint()
	return nil
}
