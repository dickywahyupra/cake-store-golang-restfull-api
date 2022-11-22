package controller

import (
	"cake-store-golang-restfull-api/helper"
	"cake-store-golang-restfull-api/src/model/request"
	"cake-store-golang-restfull-api/src/model/response"
	"cake-store-golang-restfull-api/src/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CakeController struct {
	CakeService service.CakeServiceInterface
}

func NewCakeController(cakeService service.CakeServiceInterface) CakeControllerInterface {
	return &CakeController{
		CakeService: cakeService,
	}
}

func (controller *CakeController) Create(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	cakeCreateRequest := request.CakeCreateRequest{}
	helper.ReadFromRequestBody(req, &cakeCreateRequest)

	cakeResponse := controller.CakeService.Create(req.Context(), cakeCreateRequest)
	apiResponse := response.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   cakeResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CakeController) Update(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	cakeUpdateRequest := request.CakeUpdateRequest{}
	helper.ReadFromRequestBody(req, &cakeUpdateRequest)

	cakeId := params.ByName("id")
	id, err := strconv.Atoi(cakeId)
	helper.IfError(err)

	cakeUpdateRequest.Id = id

	cakeResponse := controller.CakeService.Update(req.Context(), cakeUpdateRequest)
	apiResponse := response.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   cakeResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CakeController) Delete(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	cakeId := params.ByName("id")
	id, err := strconv.Atoi(cakeId)
	helper.IfError(err)

	controller.CakeService.Delete(req.Context(), id)
	apiResponse := response.ApiResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CakeController) FindById(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	cakeId := params.ByName("id")
	id, err := strconv.Atoi(cakeId)
	helper.IfError(err)

	cakeResponse := controller.CakeService.FindById(req.Context(), id)
	apiResponse := response.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   cakeResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CakeController) FindAll(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	cakeResponses := controller.CakeService.FindAll(req.Context())
	apiResponse := response.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   cakeResponses,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}
