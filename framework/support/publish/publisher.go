package publish

type (
	Publisher interface {
		Source() []byte
		Publish(name string) error
	}
)
