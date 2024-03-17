package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type BlogPostService struct {
	BlogRepo *repo.BlogPostRepository
}

func (service *BlogPostService) CreateBlog(blog *model.BlogPost) error {
	err := service.BlogRepo.CreateBlog(blog)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogPostService) GetAll(page, pageSize int) ([]model.BlogPost, error) {

	blogs, err := service.BlogRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (service *BlogPostService) GetById(id string) (*model.BlogPost, error) {
	blog, err := service.BlogRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("blog with id %s is not found", id)
	}
	return blog, nil
}

func (service *BlogPostService) Update(blog *model.BlogPost) error {
	err := service.BlogRepo.Update(blog)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogPostService) Delete(id string) error {
	err := service.BlogRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
