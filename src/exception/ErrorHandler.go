package exception

import (
	"cake-store-golang-restfull-api/helper"
	"cake-store-golang-restfull-api/src/model/response"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, req *http.Request, err interface{}) {
	if validationError(writer, req, err) {
		return
	}

	if notFoundError(writer, req, err) {
		return
	}

	internalServerError(writer, req, err)
}

func internalServerError(writer http.ResponseWriter, req *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	apiResponse := response.ApiResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func validationError(writer http.ResponseWriter, req *http.Request, err interface{}) bool {
	exeption, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		apiResponse := response.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad request",
			Data:   exeption.Error(),
		}

		helper.WriteToResponseBody(writer, apiResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, req *http.Request, err interface{}) bool {
	_, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		apiResponse := response.ApiResponse{
			Code:   http.StatusNotFound,
			Status: "Data not found",
			Data:   err,
		}

		helper.WriteToResponseBody(writer, apiResponse)

		return true
	} else {
		return false
	}
}
