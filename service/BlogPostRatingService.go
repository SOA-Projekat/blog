package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type BlogPostRatingService struct {
	RatingRepo *repo.BlogPostRatingRepository
}

func (service *BlogPostRatingService) CreateRating(rating *model.BlogPostRating) error {
	err := service.RatingRepo.CreateRating(rating)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogPostRatingService) GetAll(page, pageSize int) ([]model.BlogPostRating, error) {

	ratings, err := service.RatingRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (service *BlogPostRatingService) GetById(id string) (*model.BlogPostRating, error) {
	rating, err := service.RatingRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("rating with id %s is not found", id)
	}
	return rating, nil
}

func (service *BlogPostRatingService) Update(rating *model.BlogPostRating) error {
	err := service.RatingRepo.Update(rating)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogPostRatingService) Delete(id uint) error {
	err := service.RatingRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
