package responses

type (
	Paginator struct {
		Page  int         `json:"page"`
		Data  interface{} `json:"data"`
		Total int64       `json:"total"`
	}
)
