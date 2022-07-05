package delivery

import (
	"blog-api/src/helper"
	"blog-api/src/middleware"
	"blog-api/src/model"
	"blog-api/src/request"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type categoryDelivery struct {
	categoryUsecase model.CategoryUsecase
}

func NewCategoryDelivery(categoryUsecase model.CategoryUsecase) *categoryDelivery {
	return &categoryDelivery{categoryUsecase}
}

func (d *categoryDelivery) Mount(group *mux.Router, middleware middleware.AuthMiddleware) {
	group.Handle("/category", middleware.AuthRequired(http.HandlerFunc(d.GetAllCategoryHandler))).Methods("GET")
	group.Handle("/category", middleware.AuthRequired(http.HandlerFunc(d.CreateCategoryHandler))).Methods("POST")
}

func (d *categoryDelivery) CreateCategoryHandler(res http.ResponseWriter, req *http.Request) {
	request := new(request.CreateCategoryRequest)
	validator := validator.New()

	json.NewDecoder(req.Body).Decode(&request)

	if validationErr := validator.Struct(request); validationErr != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "Validation failed", Error: helper.ErrorValidation(validationErr)})
		return
	}

	err := d.categoryUsecase.CreateCategory(*request)

	if err != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "", Error: err.Error()})
		return
	}

	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "", Data: nil})
}

func (d *categoryDelivery) GetAllCategoryHandler(res http.ResponseWriter, req *http.Request) {
	categoies, err := d.categoryUsecase.GetAllCategory()
	fmt.Println(categoies)
	if err != nil {
		helper.CreateErrorResponse(res, helper.ErrorResponse{Code: http.StatusBadRequest, Message: "", Error: err.Error()})
		return
	}

	helper.CreateSuccessResponse(res, helper.SuccessResponse{Code: http.StatusOK, Message: "", Data: categoies})
}
