package model

import (
	"blog-api/src/request"
	"net/http"
	"time"
)

type (
	Post struct {
		ID         uint   `json:"id"`
		Slug       string `json:"slug"`
		Title      string `json:"title"`
		Content    string `json:"content"`
		UserId     uint   `json:"user_id"`
		User       User
		CategoryId uint `json:"category_id"`
		Category   Category
		CreatedAt  *time.Time `json:"created_at,omitempty"`
		UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	}

	PostDelivery interface {
		CreatePostHandler(res http.ResponseWriter, req *http.Request)
		GetDetailPostHandler(res http.ResponseWriter, req *http.Request)
		GetAllUserPostHandler(res http.ResponseWriter, req *http.Request)
		GetAllPostHandler(res http.ResponseWriter, req *http.Request)
		UpdatePostHandler(res http.ResponseWriter, req *http.Request)
		DeletePostHandler(res http.ResponseWriter, req *http.Request)
	}

	PostUsecase interface {
		CreatePost(request request.CreatePostRequest, userId uint) error
		GetDetailPost(id uint) (*Post, error)
		GetAllUserPost(id uint) ([]*Post, error)
		GetAllPost() ([]*Post, error)
		UpdatePost(id uint, request request.UpdatePostRequest) error
		DeletePost(id uint) error
	}

	PostRepository interface {
		InsertPost(post *Post) error
		GetPostById(id uint) (*Post, error)
		GetPostByUserId(userId uint) ([]*Post, error)
		GetAllPost() ([]*Post, error)
		UpdatePost(id uint, request request.UpdatePostRequest) error
		DeletePost(id uint) error
	}
)
