package events

type (
	BeforeSave struct {
		payload BeforeSavePayload
	}

	BeforeSavePayload struct {
		Model string
	}
)
