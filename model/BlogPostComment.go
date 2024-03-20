package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// BlogPostComment represents the structure of a comment for a blog post
type BlogPostComment struct {
	BlogPostCommentID uint      `gorm:"primaryKey" json:"blog_post_comment_id"`
	BlogID            uint      `json:"blog_id"`
	Text              string    `json:"text"`
	UserID            int       `json:"user_id"`
	CreationTime      time.Time `json:"creation_time"`
	LastUpdatedTime   time.Time `json:"last_updated_time"`
}

// NewBlogPostComment creates a new instance of BlogPostComment and returns it
func NewBlogPostComment(text string, userID int, creationTime, lastUpdatedTime time.Time) (*BlogPostComment, error) {
	if text == "" {
		return nil, errors.New("text is required")
	}
	if userID == 0 {
		return nil, errors.New("UserID is required")
	}
	if creationTime.IsZero() {
		return nil, errors.New("CreationTime is required")
	}
	if lastUpdatedTime.IsZero() {
		return nil, errors.New("LastUpdatedTime is required")
	}

	return &BlogPostComment{
		Text:            text,
		UserID:          userID,
		CreationTime:    creationTime,
		LastUpdatedTime: lastUpdatedTime,
	}, nil
}

// Value implements the driver Valuer interface
func (c BlogPostComment) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan implements the sql Scanner interface
func (c *BlogPostComment) Scan(src interface{}) error {
	if src == nil {
		*c = BlogPostComment{}
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("invalid type for BlogPostComment")
	}
	return json.Unmarshal(bytes, c)
}
