package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type BlogPostCommentService struct {
	BlogCommentRepo *repo.BlogPostCommentRepository
}

func (service *BlogPostCommentService) CreateComment(comment *model.BlogPostComment) error {
	err := service.BlogCommentRepo.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogPostCommentService) GetAll(page, pageSize int) ([]model.BlogPostComment, error) {

	comments, err := service.BlogCommentRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (service *BlogPostCommentService) GetById(id string) (*model.BlogPostComment, error) {
	comment, err := service.BlogCommentRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("comment with id %s is not found", id)
	}
	return comment, nil
}

func (service *BlogPostCommentService) Update(comment *model.BlogPostComment) error {
	err := service.BlogCommentRepo.Update(comment)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogPostCommentService) Delete(id uint) error {
	err := service.BlogCommentRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
