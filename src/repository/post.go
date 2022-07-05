package repository

import (
	"blog-api/src/model"
	"blog-api/src/request"

	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{db}
}

func (r *postRepository) InsertPost(post *model.Post) error {
	if err := r.db.Model(&model.Post{}).Create(post).Error; err != nil {
		return err
	}

	return nil
}

func (r *postRepository) GetPostById(id uint) (*model.Post, error) {
	post := new(model.Post)

	if err := r.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "slug")
	}).First(&post, id).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) GetPostByUserId(userId uint) ([]*model.Post, error) {
	var post []*model.Post

	if err := r.db.Where("user_id = ?", userId).Select("id", "slug", "title", "content", "user_id", "category_id", "created_at").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "slug")
	}).Find(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) GetAllPost() ([]*model.Post, error) {
	var post []*model.Post

	if err := r.db.Select("id", "slug", "title", "content", "user_id", "category_id", "created_at").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "slug")
	}).Find(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) UpdatePost(id uint, request request.UpdatePostRequest) error {
	if err := r.db.Model(&model.Post{}).Where("id = ?", id).Updates(model.Post{Title: request.Title, Content: request.Content, CategoryId: request.CategoryId}).Error; err != nil {
		return err
	}

	return nil
}

func (r *postRepository) DeletePost(id uint) error {
	post := new(model.Post)

	if err := r.db.First(&post, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&model.Post{}, id).Error; err != nil {
		return err
	}

	return nil
}
