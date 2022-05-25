package responses

type (
	PaginatorData interface {
	}

	Paginator struct {
		Page  int           `json:"page"`
		Data  PaginatorData `json:"data"`
		Total int64         `json:"total"`
	}
)
