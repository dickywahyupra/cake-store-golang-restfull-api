package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CakeControllerInterface interface {
	Create(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
}
