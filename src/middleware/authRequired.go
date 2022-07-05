package middleware

import (
	"blog-api/src/helper"
	"blog-api/src/repository"
	"context"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type authMiddleware struct {
	db *gorm.DB
}

type AuthMiddleware interface {
	AuthRequired(h http.Handler) http.Handler
}

func NewAuthMiddleware(db *gorm.DB) *authMiddleware {
	return &authMiddleware{db}
}

func (m *authMiddleware) AuthRequired(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		header := req.Header.Get("Authorization")
		tokenString := strings.Split(header, "Bearer ")

		if len(tokenString) < 2 {
			helper.CreateErrorResponse(res, helper.ErrorResponse{Code: 400, Message: "Error", Error: "Auth required"})
			return
		}

		claims, err := helper.DecodeTokenFromUser(tokenString[1])

		if err != nil {
			helper.CreateErrorResponse(res, helper.ErrorResponse{Code: 400, Message: "Error", Error: err.Error()})
			return
		}

		userId := uint(claims["userId"].(float64))

		userRepository := repository.NewUserRepository(m.db)

		user, errFetch := userRepository.GetUserById(userId)

		if errFetch != nil {
			helper.CreateErrorResponse(res, helper.ErrorResponse{Code: 400, Message: "Error", Error: errFetch.Error()})
			return
		}

		ctx := context.WithValue(req.Context(), "userId", user.ID)

		h.ServeHTTP(res, req.WithContext(ctx))
	})
}
