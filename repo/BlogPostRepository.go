package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type BlogPostRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogPostRepository) CreateBlog(blog *model.BlogPost) error {
	err := repo.DatabaseConnection.Create(blog).Error
	return err
}

func (repo *BlogPostRepository) GetById(id string) (*model.BlogPost, error) {
	var blog model.BlogPost
	err := repo.DatabaseConnection.First(&blog, "blog_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (repo *BlogPostRepository) GetAll(page, pageSize int) ([]model.BlogPost, error) {
	var blogs []model.BlogPost
	offset := (page - 1) * pageSize

	// Izvrši upit na bazu podataka za dohvatanje svih blogova sa straničenjem
	err := repo.DatabaseConnection.Offset(offset).Limit(pageSize).Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (repo *BlogPostRepository) Update(blog *model.BlogPost) error {
	err := repo.DatabaseConnection.Save(blog).Error
	return err
}

func (repo *BlogPostRepository) Delete(id uint) error {
	var blog model.BlogPost
	if err := repo.DatabaseConnection.Where("blog_id = ?", id).Delete(&blog).Error; err != nil {
		return err
	}
	return nil
}
