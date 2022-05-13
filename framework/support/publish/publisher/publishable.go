package publisher

type (
	Publishable struct {
		Path    string
		Content []byte
	}
)

func NewPublishable(path string, content []byte) Publishable {
	return Publishable{
		Path:    path,
		Content: content,
	}
}
