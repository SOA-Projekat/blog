package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type BlogPostCommentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogPostCommentRepository) CreateComment(comment *model.BlogPostComment) error {
	err := repo.DatabaseConnection.Create(comment).Error
	return err
}

func (repo *BlogPostCommentRepository) GetById(id string) (*model.BlogPostComment, error) {
	var comment model.BlogPostComment
	err := repo.DatabaseConnection.First(&comment, "comment_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (repo *BlogPostCommentRepository) GetAll(page, pageSize int) ([]model.BlogPostComment, error) {
	var comments []model.BlogPostComment
	offset := (page - 1) * pageSize

	// Izvrši upit na bazu podataka za dohvatanje svih blogova sa straničenjem
	err := repo.DatabaseConnection.Offset(offset).Limit(pageSize).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (repo *BlogPostCommentRepository) Update(comment *model.BlogPostComment) error {
	err := repo.DatabaseConnection.Save(comment).Error
	return err
}

func (repo *BlogPostCommentRepository) Delete(id uint) error {
	var comment model.BlogPostComment
	if err := repo.DatabaseConnection.Where("comment_id = ?", id).Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
