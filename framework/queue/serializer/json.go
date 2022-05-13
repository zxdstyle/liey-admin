package serializer

import "encoding/json"

type JsonSerializer struct {
}

func (j JsonSerializer) Serialize(payload interface{}) ([]byte, error) {
	return json.Marshal(payload)
}

func (j JsonSerializer) UnSerialize(payload []byte, pointer interface{}) error {
	return json.Unmarshal(payload, pointer)
}
