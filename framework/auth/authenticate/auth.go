package authenticate

type (
	Authenticate interface {
		GetKey() uint
		GuardName() string
	}

	DefaultAuthenticate struct {
		ID    uint
		Guard string
	}
)

func (auth DefaultAuthenticate) GetKey() uint {
	return auth.ID
}

func (auth DefaultAuthenticate) GuardName() string {
	return auth.Guard
}

func NewDefaultAuthenticate(id uint, guard string) Authenticate {
	return DefaultAuthenticate{
		ID:    id,
		Guard: guard,
	}
}
