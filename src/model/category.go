package model

import (
	"blog-api/src/request"
	"net/http"
	"time"
)

type (
	Category struct {
		ID        uint       `json:"id"`
		Name      string     `json:"name"`
		Slug      string     `json:"slug"`
		Posts     []Post     `json:"post,omitempty"`
		CreatedAt *time.Time `json:"created_at,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}

	CategoryDelivery interface {
		CreateCategoryHandler(res http.ResponseWriter, req *http.Request)
	}

	CategoryUsecase interface {
		CreateCategory(request request.CreateCategoryRequest) error
		GetAllCategory() ([]*Category, error)
	}

	CategoryRepository interface {
		InsertCategory(category *Category) error
		FindCategoryById(id uint) (*Category, error)
		GetAllCategory() ([]*Category, error)
	}
)
