package authenticate

type (
	Authenticate interface {
		GetKey() uint
		Guard() string
	}
)
