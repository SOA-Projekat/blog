package service

import (
	"database-example/model"
	"database-example/repo"
	"time"
)

type BlogPostRatingService struct {
	RatingRepo *repo.BlogPostRatingRepository
}

// AddRating dodaje novu ocenu za dati blog post.
func (service *BlogPostRatingService) AddRating(blogID uint, userID int, isPositive bool) error {
	// Kreiranje nove ocene
	rating := model.NewBlogPostRating(blogID, userID, time.Now(), isPositive)

	// Dodavanje ocene pomoću repozitorijuma
	err := service.RatingRepo.AddRating(blogID, rating)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRating briše ocenu za dati blog post na osnovu korisničkog ID-ja.
func (service *BlogPostRatingService) DeleteRating(blogID uint, userID int) error {
	// Brisanje ocene pomoću repozitorijuma
	err := service.RatingRepo.DeleteRating(blogID, userID)
	if err != nil {
		return err
	}

	return nil
}
