package helper

import (
	"encoding/json"
	"net/http"
)

type (
	SuccessResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	ErrorResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
	}
)

func CreateSuccessResponse(res http.ResponseWriter, data SuccessResponse) {
	resJson, _ := json.Marshal(data)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(data.Code)
	res.Write(resJson)
}

func CreateErrorResponse(res http.ResponseWriter, data ErrorResponse) {
	resJson, _ := json.Marshal(data)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(data.Code)
	res.Write(resJson)
}
