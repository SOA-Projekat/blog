package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type BlogPostRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogPostRepository) CreateBlog(blog *model.Blog) error {
	err := repo.DatabaseConnection.Create(blog).Error
	return err
}

func (repo *BlogPostRepository) GetById(id string) (*model.Blog, error) {
	var blog model.Blog
	err := repo.DatabaseConnection.First(&blog, "blog_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (repo *BlogPostRepository) GetAll(page, pageSize int) ([]model.Blog, error) {
	var blogs []model.Blog
	offset := (page - 1) * pageSize

	// Izvrši upit na bazu podataka za dohvatanje svih blogova sa straničenjem
	err := repo.DatabaseConnection.Offset(offset).Limit(pageSize).Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (repo *BlogPostRepository) Update(blog *model.Blog) error {
	err := repo.DatabaseConnection.Save(blog).Error
	return err
}

func (repo *BlogPostRepository) Delete(id string) error {
	var blog model.Blog
	if err := repo.DatabaseConnection.Where("blog_id = ?", id).Delete(&blog).Error; err != nil {
		return err
	}
	return nil
}
