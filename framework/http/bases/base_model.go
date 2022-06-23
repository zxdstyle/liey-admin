package bases

import "github.com/gogf/gf/v2/os/gtime"

type (
	Model struct {
		ID        uint        `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
		CreatedAt *gtime.Time `gorm:"not null;index:idx_created_at;comment:创建时间" json:"created_at"`
		UpdatedAt *gtime.Time `gorm:"comment:更新时间" json:"updated_at"`
	}

	CreateOnlyModel struct {
		ID        uint        `gorm:"primaryKey;autoIncrement；comment:主键" json:"id"`
		CreatedAt *gtime.Time `gorm:"not null;index:idx_created_at;comment:创建时间" json:"created_at"`
	}

	OnlyKeyModel struct {
		ID uint `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	}
)

func (m Model) GetKey() uint {
	return m.ID
}

func (m *Model) SetKey(id uint) {
	m.ID = id
}

func (m Model) GetCreatedAt() *gtime.Time {
	return m.CreatedAt
}

func (m *Model) SetCreatedAt(ct *gtime.Time) {
	m.CreatedAt = ct
}

func (m Model) GetUpdatedAt() *gtime.Time {
	return m.UpdatedAt
}

func (m *Model) SetUpdatedAt(ut *gtime.Time) {
	m.UpdatedAt = ut
}

func (c CreateOnlyModel) GetKey() uint {
	return c.ID
}

func (c *CreateOnlyModel) SetKey(id uint) {
	c.ID = id
}

func (c CreateOnlyModel) GetCreatedAt() *gtime.Time {
	return c.CreatedAt
}

func (c *CreateOnlyModel) SetCreatedAt(ct *gtime.Time) {
	c.CreatedAt = ct
}

func (c CreateOnlyModel) GetUpdatedAt() *gtime.Time {
	return nil
}

func (c *CreateOnlyModel) SetUpdatedAt(ut *gtime.Time) {

}

func (c OnlyKeyModel) GetKey() uint {
	return c.ID
}

func (c *OnlyKeyModel) SetKey(id uint) {
	c.ID = id
}

func (c OnlyKeyModel) GetCreatedAt() *gtime.Time {
	return nil
}

func (c *OnlyKeyModel) SetCreatedAt(ct *gtime.Time) {
}

func (c OnlyKeyModel) GetUpdatedAt() *gtime.Time {
	return nil
}

func (c *OnlyKeyModel) SetUpdatedAt(ut *gtime.Time) {
}
