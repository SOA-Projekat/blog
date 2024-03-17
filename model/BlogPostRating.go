package model

import (
	"time"
)

// BlogPostRating predstavlja ocenu ili rejting datog blog posta.
type BlogPostRating struct {
	UserID       int       `json:"userId"`
	CreationTime time.Time `json:"creationTime"`
	IsPositive   bool      `json:"isPositive"`
}

// NewBlogPostRating kreira novu instancu BlogPostRating i vraća je.
func NewBlogPostRating(userID int, creationTime time.Time, isPositive bool) *BlogPostRating {
	return &BlogPostRating{
		UserID:       userID,
		CreationTime: creationTime,
		IsPositive:   isPositive,
	}
}
