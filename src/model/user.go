package model

import (
	"blog-api/src/request"
	"net/http"
	"time"
)

type (
	User struct {
		ID        uint       `json:"id"`
		Username  string     `json:"username"`
		Email     string     `json:"email,omitempty"`
		Password  string     `json:"password,omitempty"`
		Posts     []Post     `json:"post,omitempty"`
		CreatedAt *time.Time `json:"created_at,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}

	UserDeliver interface {
		UserSignupHandler(res http.ResponseWriter, req *http.Request)
		UserLoginHandler(res http.ResponseWriter, req *http.Request)
		GetUserHandler(res http.ResponseWriter, req *http.Request)
		UpdateUsernameHandler(res http.ResponseWriter, req *http.Request)
	}

	UserUsecase interface {
		GetUser(userId uint) (*User, error)
		UserSignup(request request.UserSignupRequest) error
		UserLogin(request request.UserLoginRequest) (string, error)
		UpdateUsername(userId uint, request request.UpdateUsernameRequest) error
	}

	UserRepository interface {
		InsertUser(user *User) error
		GetUserById(userId uint) (*User, error)
		GetUserByEmail(email string) (*User, error)
		GetCountUserByEmail(email string) (int64, error)
		UpdateUser(userId uint, username string) error
	}
)
