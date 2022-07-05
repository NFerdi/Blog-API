package repository

import (
	"blog-api/src/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) InsertUser(user *model.User) error {
	if err := r.db.Model(&model.User{}).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserById(userId uint) (*model.User, error) {
	user := new(model.User)
	if err := r.db.Model(&model.User{}).Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "slug", "title", "content", "category_id", "user_id").Preload("Category")
	}).First(user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	user := new(model.User)
	if err := r.db.Model(&model.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetCountUserByEmail(email string) (int64, error) {
	count := int64(0)
	if err := r.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *userRepository) UpdateUser(userId uint, username string) error {
	if err := r.db.Model(&model.User{}).Where("id = ?", userId).Update("username", username).Error; err != nil {
		return err
	}

	return nil
}
