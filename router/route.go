package router

import (
	"cake-store-golang-restfull-api/src/controller"

	"github.com/julienschmidt/httprouter"
)

func GetRoute(cakeController controller.CakeControllerInterface) *httprouter.Router {
	router := httprouter.New()
	prefix := "/api"

	router.GET(prefix+"/cakes", cakeController.FindAll)
	router.GET(prefix+"/cakes/:id", cakeController.FindById)
	router.POST(prefix+"/cakes", cakeController.Create)
	router.PUT(prefix+"/cakes/:id", cakeController.Update)
	router.DELETE(prefix+"/cakes/:id", cakeController.Delete)

	return router
}
