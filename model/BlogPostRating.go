package model

import "time"

type BlogPostRating struct {
	RatingID     uint      `gorm:"primaryKey" json:"rating_id"`
	IsPositive   bool      `json:"is_positive"`
	CreationTime time.Time `json:"creation_time"`
	UserId       int       `json:"user_id"`
}
