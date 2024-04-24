package model

import (
	//"database/sql/driver"
	//"encoding/json"
	//"errors"
	"time"
)

type BlogPostComment struct {
	CommentID       uint      `gorm:"primaryKey" json:"comment_id"`
	Text            string    `json:"text"`
	UserId          int       `json:"user_id"`
	Username        string    `json:"username,omitempty"`
	CreationTime    time.Time `json:"creation_time"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}
