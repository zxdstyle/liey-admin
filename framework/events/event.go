package events

type (
	Event interface {
		Payload() interface{}
	}
)
