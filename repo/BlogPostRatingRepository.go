package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type BlogPostRatingRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogPostRatingRepository) CreateRating(rating *model.BlogPostRating) error {
	err := repo.DatabaseConnection.Create(rating).Error
	return err
}

func (repo *BlogPostRatingRepository) GetById(id string) (*model.BlogPostRating, error) {
	var rating model.BlogPostRating
	err := repo.DatabaseConnection.First(&rating, "rating_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

func (repo *BlogPostRatingRepository) GetAll(page, pageSize int) ([]model.BlogPostRating, error) {
	var ratings []model.BlogPostRating
	offset := (page - 1) * pageSize

	// Izvrši upit na bazu podataka za dohvatanje svih blogova sa straničenjem
	err := repo.DatabaseConnection.Offset(offset).Limit(pageSize).Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (repo *BlogPostRatingRepository) Update(rating *model.BlogPostRating) error {
	err := repo.DatabaseConnection.Save(rating).Error
	return err
}

func (repo *BlogPostRatingRepository) Delete(id uint) error {
	var rating model.BlogPostRating
	if err := repo.DatabaseConnection.Where("rating_id = ?", id).Delete(&rating).Error; err != nil {
		return err
	}
	return nil
}
