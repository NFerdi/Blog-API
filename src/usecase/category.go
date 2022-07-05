package usecase

import (
	"blog-api/src/model"
	"blog-api/src/request"

	"github.com/gosimple/slug"
)

type categoryUsecase struct {
	categoryRepository model.CategoryRepository
}

func NewCategoryUsecase(categoryRepository model.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{categoryRepository}
}

func (uc *categoryUsecase) CreateCategory(request request.CreateCategoryRequest) error {
	slug := slug.Make(request.Name)
	category := &model.Category{Name: request.Name, Slug: slug}

	if err := uc.categoryRepository.InsertCategory(category); err != nil {
		return err
	}

	return nil
}

func (uc *categoryUsecase) GetAllCategory() ([]*model.Category, error) {
	categories, err := uc.categoryRepository.GetAllCategory()

	if err != nil {
		return nil, err
	}

	return categories, nil
}
