package serializer

type Serializer interface {
	Serialize(payload interface{}) ([]byte, error)
	UnSerialize(payload []byte, pointer interface{}) error
}
