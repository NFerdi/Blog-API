package usecase

import (
	"blog-api/src/helper"
	"blog-api/src/model"
	"blog-api/src/request"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository model.UserRepository
}

func NewUserUsecase(userRepository model.UserRepository) *userUsecase {
	return &userUsecase{userRepository}
}

func (uc *userUsecase) UserSignup(request request.UserSignupRequest) error {
	c, err := uc.userRepository.GetCountUserByEmail(request.Email)

	if err != nil {
		return err
	}

	if c > 0 {
		return errors.New("email already register")
	}

	hash, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return errHash
	}

	request.Password = string(hash)
	user := &model.User{Username: request.Username, Email: request.Email, Password: request.Password}
	errInsert := uc.userRepository.InsertUser(user)

	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (uc *userUsecase) UserLogin(request request.UserLoginRequest) (string, error) {
	c, err := uc.userRepository.GetCountUserByEmail(request.Email)
	token := string("")

	if err != nil {
		return token, err
	}

	if c != 1 {
		return token, errors.New("email not register")
	}

	user, errFetch := uc.userRepository.GetUserByEmail(request.Email)

	if errFetch != nil {
		return token, errFetch
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if errCompare != nil {
		return token, errors.New("password does not match")
	}

	token, errToken := helper.GenerateTokenFromUser(user.ID)

	if errToken != nil {
		return token, errToken
	}

	return token, nil
}

func (uc *userUsecase) GetUser(userId uint) (*model.User, error) {
	user, err := uc.userRepository.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUsecase) UpdateUsername(userId uint, request request.UpdateUsernameRequest) error {
	if err := uc.userRepository.UpdateUser(userId, request.Username); err != nil {
		return err
	}

	return nil
}
