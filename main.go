package main

import (
	"cake-store-golang-restfull-api/database"
	"cake-store-golang-restfull-api/helper"
	"cake-store-golang-restfull-api/router"
	"cake-store-golang-restfull-api/src/controller"
	"cake-store-golang-restfull-api/src/middleware"
	"cake-store-golang-restfull-api/src/repository"
	"cake-store-golang-restfull-api/src/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
)

func main() {
	db := database.MysqlConnect()
	database.Migration(db)
	validate := validator.New()

	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, db, validate)
	cakeController := controller.NewCakeController(cakeService)

	router := router.GetRoute(cakeController)

	handler := cors.AllowAll().Handler(middleware.NewAuthMiddleware(router))

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.IfError(err)
}
