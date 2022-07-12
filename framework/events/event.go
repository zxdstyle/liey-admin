package events

type (
	Event[V any] interface {
		Payload() V
	}

	NamedEvent interface {
		Name() string
	}
)
