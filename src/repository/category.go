package repository

import (
	"blog-api/src/model"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) InsertCategory(category *model.Category) error {
	if err := r.db.Model(&model.Category{}).Create(category).Error; err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) FindCategoryById(id uint) (*model.Category, error) {
	category := new(model.Category)

	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) GetAllCategory() ([]*model.Category, error) {
	var categories []*model.Category

	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
