package delivery

import (
	"blog-api/src/helper"
	"blog-api/src/middleware"
	"blog-api/src/model"
	"blog-api/src/request"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type postDelivery struct {
	postUsecase model.PostUsecase
}

func NewPostDelivery(postUsecase model.PostUsecase) *postDelivery {
	return &postDelivery{postUsecase}
}

func (d *postDelivery) Mount(group *mux.Router, middleware middleware.AuthMiddleware) {
	group.Handle("/post", middleware.AuthRequired(http.HandlerFunc(d.CreatePostHandler))).Methods("POST")
	group.Handle("/post", middleware.AuthRequired(http.HandlerFunc(d.GetAllPostHandler))).Methods("GET")
	group.Handle("/post/user", middleware.AuthRequired(http.HandlerFunc(d.GetAllUserPostHandler))).Methods("GET")
	group.Handle("/post/{id}", middleware.AuthRequired(http.HandlerFunc(d.GetDetailPostHandler))).Methods("GET")
	group.Handle("/post/{id}", middleware.AuthRequired(http.HandlerFunc(d.UpdatePostHandler))).Methods("PATCH")
	group.Handle("/post/{id}", middleware.AuthRequired(http.HandlerFunc(d.DeletePostHandler))).Methods("DELETE")
}

func (d *postDelivery) CreatePostHandler(res http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value("userId").(uint)
	request := new(request.CreatePostRequest)
	validator := validator.New()

	json.NewDecoder(req.Body).Decode(&request)

	if validationErr := validator.Struct(request); validationErr != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Validation failed", Error: helper.ErrorValidation(validationErr)})
		return
	}

	if err := d.postUsecase.CreatePost(*request, userId); err != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Error creating user post", Error: err.Error()})
		return
	}
	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "Successfully created a user post", Data: nil})
}

func (d *postDelivery) GetDetailPostHandler(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.ParseUint(string(mux.Vars(req)["id"]), 10, 32)

	post, err := d.postUsecase.GetDetailPost(uint(id))

	if err != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "An error occurred while retrieving post details", Error: err.Error()})
		return
	}
	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "Successfully retrieved post details", Data: post})
}

func (d *postDelivery) GetAllUserPostHandler(res http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value("userId").(uint)

	post, err := d.postUsecase.GetAllUserPost(userId)

	if err != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Error retrieving user posts", Error: err.Error()})
		return
	}
	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "Successfully retrieved user posts", Data: post})
}

func (d *postDelivery) GetAllPostHandler(res http.ResponseWriter, req *http.Request) {
	post, err := d.postUsecase.GetAllPost()

	if err != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Error retrieving all posts", Error: err.Error()})
		return
	}
	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "Successfully retrieved all posts", Data: post})
}

func (d *postDelivery) UpdatePostHandler(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.ParseUint(string(mux.Vars(req)["id"]), 10, 32)
	request := new(request.UpdatePostRequest)
	validator := validator.New()

	json.NewDecoder(req.Body).Decode(&request)

	if validationErr := validator.Struct(request); validationErr != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Validation failed", Error: helper.ErrorValidation(validationErr)})
		return
	}

	err := d.postUsecase.UpdatePost(uint(id), *request)

	if err != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Error while updating post", Error: err.Error()})
		return
	}

	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "Successfully updated post", Data: nil})
}

func (d *postDelivery) DeletePostHandler(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.ParseUint(string(mux.Vars(req)["id"]), 10, 32)

	err := d.postUsecase.DeletePost(uint(id))

	if err != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Error deleting post", Error: err.Error()})
		return
	}

	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "Successfully deleted post", Data: nil})
}
