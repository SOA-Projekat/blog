package model

import (
	"errors"
	"time"
)

// BlogPostComment represents the structure of a comment for a blog post
type BlogPostComment struct {
	Text            string    `json:"text"`
	UserID          int       `json:"user_id"`
	CreationTime    time.Time `json:"creation_time"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
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
