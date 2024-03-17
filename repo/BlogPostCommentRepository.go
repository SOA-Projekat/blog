package repo

import (
	"database-example/model"
	"time"

	"gorm.io/gorm"
)

type BlogPostCommentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogPostCommentRepository) AddComment(blogID int, comment *model.BlogPostComment) error {
	// Postavite vreme kreiranja ako nije postavljeno
	if comment.CreationTime.IsZero() {
		comment.CreationTime = time.Now()
	}

	// Dodajte komentar u bazu podataka
	if err := repo.DatabaseConnection.Create(comment).Error; err != nil {
		return err
	}

	return nil
}

func (repo *BlogPostCommentRepository) UpdateComment(commentID int, updatedComment *model.BlogPostComment) error {
	// Pronađite komentar sa datim ID-om
	var existingComment model.BlogPostComment
	if err := repo.DatabaseConnection.First(&existingComment, commentID).Error; err != nil {
		return err
	}

	// Ažurirajte polja komentara
	existingComment.Text = updatedComment.Text
	existingComment.LastUpdatedTime = time.Now() // Postavite vreme poslednjeg ažuriranja na trenutno vreme

	// Ažurirajte komentar u bazi podataka
	if err := repo.DatabaseConnection.Save(&existingComment).Error; err != nil {
		return err
	}

	return nil
}

func (repo *BlogPostCommentRepository) DeleteComment(commentID int) error {
	// Pronađite komentar sa datim ID-om
	var existingComment model.BlogPostComment
	if err := repo.DatabaseConnection.First(&existingComment, commentID).Error; err != nil {
		return err
	}

	// Obrišite komentar iz baze podataka
	if err := repo.DatabaseConnection.Delete(&existingComment).Error; err != nil {
		return err
	}

	return nil
}
