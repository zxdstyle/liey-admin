package bases

import "github.com/gogf/gf/v2/os/gtime"

type (
	Model struct {
		ID        uint        `json:"id"`
		CreatedAt *gtime.Time `json:"created_at"`
		UpdatedAt *gtime.Time `json:"updated_at"`
	}

	CreateOnlyModel struct {
		ID        uint        `json:"id"`
		CreatedAt *gtime.Time `json:"created_at"`
	}
)

func (m Model) GetKey() uint {
	return m.ID
}

func (m Model) GetCreatedAt() *gtime.Time {
	return m.CreatedAt
}

func (m Model) GetUpdatedAt() *gtime.Time {
	return m.UpdatedAt
}

func (c CreateOnlyModel) GetKey() uint {
	return c.ID
}

func (c CreateOnlyModel) GetCreatedAt() *gtime.Time {
	return c.CreatedAt
}
