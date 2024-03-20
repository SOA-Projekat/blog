package model

import (
	"time"
)

// BlogPostRating predstavlja ocenu ili rejting datog blog posta.
type BlogPostRating struct {
	BlogPostRatingID uint      `gorm:"primaryKey" json:"blog_post_rating_id"`
	BlogID           uint      `json:"blog_id"` // Dodajte polje BlogID za povezivanje sa odgovarajućim blogom
	UserID           int       `json:"user_id"`
	CreationTime     time.Time `json:"creation_time"`
	IsPositive       bool      `json:"is_positive"`
}

// NewBlogPostRating kreira novu instancu BlogPostRating i vraća je.
func NewBlogPostRating(blogID uint, userID int, creationTime time.Time, isPositive bool) *BlogPostRating {
	return &BlogPostRating{

		BlogID:       blogID, // Dodajte polje BlogID u konstruktoru
		UserID:       userID,
		CreationTime: creationTime,
		IsPositive:   isPositive,
	}
}
