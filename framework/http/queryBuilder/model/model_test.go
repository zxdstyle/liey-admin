package model

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gtime"
	"testing"
)

type (
	Model struct {
		ID        uint        `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
		CreatedAt *gtime.Time `gorm:"not null;index:idx_created_at;comment:创建时间" json:"created_at"`
		UpdatedAt *gtime.Time `gorm:"comment:更新时间" json:"updated_at"`
	}

	TestModel struct {
		Model
		Username string
		Order    string `gorm:"column:sort_num" json:"order"`
		IsActive uint   `json:"status,omitempty"`
		//Ignore   uint   `json:"-"`
		Content string `gorm:"-"`
	}
)

func TestNew(t *testing.T) {
	mo := New(TestModel{})

	fields := garray.NewStrArrayFrom([]string{"id", "username", "sort_num", "status", "created_at", "updated_at"})

	if mo.Fields.Len() != fields.Len() {
		t.Fatal("字段数量不一致")
	}

	fields.Iterator(func(k int, v string) bool {
		if !mo.Fields.Contains(v) {
			t.Fatalf("缺少字段: %s", v)
			return false
		}
		return true
	})
}
