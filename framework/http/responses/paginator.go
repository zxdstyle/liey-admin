package responses

type (
	Meta struct {
		Page  int   `json:"page"`
		Total int64 `json:"total"`
	}

	Paginator struct {
		Data interface{} `json:"data"`
		Meta Meta        `json:"meta"`
	}
)
