package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type BlogPostRatingRepository struct {
	DatabaseConnection *gorm.DB
}

// AddRating dodaje ocenu za dati blog post.
func (repo *BlogPostRatingRepository) AddRating(blogID int, rating *model.BlogPostRating) error {
	return repo.DatabaseConnection.Create(rating).Error
}

// DeleteRating briše ocenu za dati blog post na osnovu korisničkog ID-ja i ID-ja blog posta.
func (repo *BlogPostRatingRepository) DeleteRating(blogID int, userID int) error {
	return repo.DatabaseConnection.Where("blog_id = ? AND user_id = ?", blogID, userID).Delete(&model.BlogPostRating{}).Error
}
