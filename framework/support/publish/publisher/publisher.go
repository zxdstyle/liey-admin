package publisher

type (
	Publisher interface {
		Sources() []Publishable
		Publish() error
	}
)
