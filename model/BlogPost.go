package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type BlogPostStatus int

const (
	DRAFT BlogPostStatus = iota
	PUBLISHED
	CLOSED
	ACTIVE
	FAMOUS
)

type BlogPost struct {
	BlogID       uint      `gorm:"primaryKey" json:"blog_id"`
	AuthorID     int       `json:"author_id"`
	TourID       int       `json:"tour_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creation_date"`
	ImageURLs    string    `json:"imageUrl"`
	//Comments     []BlogPostComment `json:"blog_post_comments" gorm:"foreignKey:comment_id"`
	//Ratings      []BlogPostRating  `json:"blog_post_ratings" gorm:"foreignKey:rating_id"`
	Status BlogPostStatus `json:"status"`
}

func NewBlogPost(blogID, authorID, tourID int, title, description string, creationDate time.Time, imageURLs string, status BlogPostStatus) (*BlogPost, error) {
	if authorID == 0 {
		return nil, errors.New("field required: AuthorID")
	}
	if title == "" {
		return nil, errors.New("invalid Title")
	}
	if description == "" {
		return nil, errors.New("invalid Description")
	}
	if creationDate.IsZero() {
		return nil, errors.New("invalid Creation Date")
	}
	if status != DRAFT && status != PUBLISHED && status != CLOSED && status != ACTIVE && status != FAMOUS {
		return nil, errors.New("invalid Post Status")
	}

	return &BlogPost{
		BlogID:       uint(blogID),
		AuthorID:     authorID,
		TourID:       tourID,
		Title:        title,
		Description:  description,
		CreationDate: creationDate,
		ImageURLs:    imageURLs,
		Status:       status,
	}, nil
}

type StringArray []string

func (strArray StringArray) Value() (driver.Value, error) {
	return json.Marshal(strArray)
}

func (str *StringArray) Scan(result interface{}) error {
	if result == nil {
		*str = nil
		return nil
	}
	m, n := result.([]byte)
	if !n {
		return errors.New("process of type asserting to []byte has failed")
	}
	return json.Unmarshal(m, str)
}
