package service

import (
	"database-example/model"
	"database-example/repo"
	"time"
)

type BlogPostCommentService struct {
	CommentRepo *repo.BlogPostCommentRepository
}

func (service *BlogPostCommentService) AddComment(blogID uint, text string, userID int) error {
	// Kreirajte novi komentar
	comment, err := model.NewBlogPostComment(text, userID, time.Now(), time.Now())
	if err != nil {
		return err
	}

	// Dodajte komentar pomoću repozitorijuma
	err = service.CommentRepo.AddComment(blogID, comment)
	if err != nil {
		return err
	}

	return nil
}

func (service *BlogPostCommentService) UpdateComment(commentID int, text string) error {
	// Kreirajte ažurirani komentar
	updatedComment, err := model.NewBlogPostComment(text, 0, time.Time{}, time.Now())
	if err != nil {
		return err
	}

	// Ažurirajte komentar pomoću repozitorijuma
	err = service.CommentRepo.UpdateComment(commentID, updatedComment)
	if err != nil {
		return err
	}

	return nil
}

func (service *BlogPostCommentService) DeleteComment(commentID int) error {
	// Obrišite komentar pomoću repozitorijuma
	err := service.CommentRepo.DeleteComment(commentID)
	if err != nil {
		return err
	}

	return nil
}
