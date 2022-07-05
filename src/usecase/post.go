package usecase

import (
	"blog-api/src/model"
	"blog-api/src/request"
	"errors"

	"github.com/gosimple/slug"
)

type postUsecase struct {
	postRepository     model.PostRepository
	categoryRepository model.CategoryRepository
}

func NewPostUsecase(postRepository model.PostRepository, categoryRepository model.CategoryRepository) *postUsecase {
	return &postUsecase{postRepository, categoryRepository}
}

func (uc *postUsecase) CreatePost(request request.CreatePostRequest, userId uint) error {
	slug := slug.Make(request.Title)
	post := &model.Post{Slug: slug, Title: request.Title, Content: request.Content, UserId: userId, CategoryId: request.CategoryId}

	err := uc.postRepository.InsertPost(post)

	if err != nil {
		return err
	}

	return nil
}

func (uc *postUsecase) GetDetailPost(id uint) (*model.Post, error) {
	post, err := uc.postRepository.GetPostById(id)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (uc *postUsecase) GetAllUserPost(userId uint) ([]*model.Post, error) {
	post, err := uc.postRepository.GetPostByUserId(userId)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (uc *postUsecase) GetAllPost() ([]*model.Post, error) {
	post, err := uc.postRepository.GetAllPost()

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (uc *postUsecase) UpdatePost(id uint, request request.UpdatePostRequest) error {
	post, err := uc.postRepository.GetPostById(id)

	if err != nil {
		return err
	}

	if request.Title == "" {
		request.Title = post.Title
	}
	if request.Content == "" {
		request.Content = post.Content
	}
	if request.CategoryId == 0 {
		request.CategoryId = post.CategoryId
	}

	_, errCategory := uc.categoryRepository.FindCategoryById(request.CategoryId)

	if errCategory != nil {
		return errors.New("category could not be found")
	}

	errUpdate := uc.postRepository.UpdatePost(id, request)

	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (uc *postUsecase) DeletePost(id uint) error {
	if err := uc.postRepository.DeletePost(id); err != nil {
		return err
	}

	return nil
}
