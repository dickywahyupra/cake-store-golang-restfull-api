package middleware

import (
	"cake-store-golang-restfull-api/helper"
	"cake-store-golang-restfull-api/src/model/response"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("X-API-Key")

	if header == "ROOT" {
		middleware.Handler.ServeHTTP(writer, req)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		apiResponse := response.ApiResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writer, apiResponse)
	}
}
