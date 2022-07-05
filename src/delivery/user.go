package delivery

import (
	"blog-api/src/helper"
	"blog-api/src/middleware"
	"blog-api/src/model"
	"blog-api/src/request"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type userDelivery struct {
	userUsecase model.UserUsecase
}

func NewUserDelivery(userUsecase model.UserUsecase) *userDelivery {
	return &userDelivery{userUsecase}
}

func (d *userDelivery) Mount(group *mux.Router, middleware middleware.AuthMiddleware) {
	group.HandleFunc("/auth/signup", d.UserSignupHandler).Methods("POST")
	group.HandleFunc("/auth/login", d.UserLoginHandler).Methods("POST")
	group.Handle("/user", middleware.AuthRequired(http.HandlerFunc(d.UpdateUsernameHandler))).Methods("PATCH")
	group.Handle("/user", middleware.AuthRequired(http.HandlerFunc(d.GetUserHandler))).Methods("GET")
}

func (d *userDelivery) UserSignupHandler(res http.ResponseWriter, req *http.Request) {
	request := new(request.UserSignupRequest)
	validate := validator.New()

	json.NewDecoder(req.Body).Decode(&request)

	validationErr := validate.Struct(request)

	if validationErr != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Validation failed", Error: helper.ErrorValidation(validationErr)})
		return
	}

	errSignup := d.userUsecase.UserSignup(*request)

	if errSignup != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Error when registering", Error: errSignup.Error()})
		return
	}
	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "Successfully signed up", Data: nil})
}

func (d *userDelivery) UserLoginHandler(res http.ResponseWriter, req *http.Request) {
	request := new(request.UserLoginRequest)
	validate := validator.New()

	json.NewDecoder(req.Body).Decode(&request)

	validationErr := validate.Struct(request)

	if validationErr != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Validation failed", Error: helper.ErrorValidation(validationErr)})
		return
	}

	token, errLogin := d.userUsecase.UserLogin(*request)

	if errLogin != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Error when logging in", Error: errLogin.Error()})
		return
	}
	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "Successfully logged in", Data: token})
}

func (d *userDelivery) GetUserHandler(res http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value("userId").(uint)

	user, err := d.userUsecase.GetUser(userId)

	if err != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Error while retrieving user data", Error: err.Error()})
		return
	}
	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "Successfully retrieved user data", Data: user})
}

func (d *userDelivery) UpdateUsernameHandler(res http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value("userId").(uint)
	request := new(request.UpdateUsernameRequest)
	validate := validator.New()

	json.NewDecoder(req.Body).Decode(&request)

	validationErr := validate.Struct(request)

	if validationErr != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Validation failed", Error: helper.ErrorValidation(validationErr)})
		return
	}

	if errUpdate := d.userUsecase.UpdateUsername(userId, *request); errUpdate != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "An error occurred while updating user data", Error: errUpdate.Error()})
		return
	}
	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "Successfully updated user data", Data: nil})
}
