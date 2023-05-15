package models

import (
	"github.com/practice/models/user"
	"github.com/practice/pkg/database"
	"github.com/spf13/cast"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

// GetStringID 获取ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel user.User) {
	database.DB.Where("phone = ?", loginID).Or("email = ?", loginID).Or("name = ?", loginID).First(&userModel)
	return
}
