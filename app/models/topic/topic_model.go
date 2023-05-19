//Package topic 模型
package topic

import (
    "github.com/practice/app/models"
    "github.com/practice/app/models/user"
    "github.com/practice/pkg/database"
)

type Topic struct {
    models.BaseModel

    Title      string `json:"title,omitempty"`
    Body       string `json:"body,omitempty"`
    UserID     string `json:"user_id,omitempty"`
    CategoryID string `json:"category_id,omitempty"`

    User user.User `json:"user"`

    models.CommonTimestampsField
}

func (topic *Topic) Create() {
    database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
    result := database.DB.Save(&topic)
    return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
    result := database.DB.Delete(&topic)
    return result.RowsAffected
}