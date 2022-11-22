package test

import (
	"cake-store-golang-restfull-api/helper"
	"cake-store-golang-restfull-api/router"
	"cake-store-golang-restfull-api/src/controller"
	"cake-store-golang-restfull-api/src/middleware"
	"cake-store-golang-restfull-api/src/model/domain"
	"cake-store-golang-restfull-api/src/repository"
	"cake-store-golang-restfull-api/src/service"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cake-store")
	helper.IfError(err)

	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, db, validate)
	cakeController := controller.NewCakeController(cakeService)

	router := router.GetRoute(cakeController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCake(db *sql.DB) {
	db.Exec("TRUNCATE cakes")
}

func initCake(db *sql.DB) *domain.Cake {
	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cake := cakeRepository.Save(context.Background(), tx, domain.Cake{
		Title:       "Lemon cheesecake",
		Description: "A cheesecake made of lemon",
		Rating:      7.3,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	})
	tx.Commit()

	return &cake
}

func transformResponse(response *http.Response) map[string]interface{} {
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	return responseBody
}

var baseUrl string = "http://localhost/api"

func TestCreateCakeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"title": "Lemon cheesecake",
		"description": "A cheesecake made of lemon",
		"rating": 7.3,
		"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
	}`)
	request := httptest.NewRequest(http.MethodPost, baseUrl+"/cakes", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "ROOT")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	responseBody := transformResponse(response)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	data := responseBody["data"].(map[string]interface{})
	assert.Equal(t, "Lemon cheesecake", data["title"])
	assert.Equal(t, "A cheesecake made of lemon", data["description"])
	assert.Equal(t, 7.3, data["rating"])
	assert.Equal(t, "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", data["image"])
}

func TestCreateCakeFailed(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"title":""}`)
	request := httptest.NewRequest(http.MethodPost, baseUrl+"/cakes", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "ROOT")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad request", responseBody["status"])
}

func TestUpdateCakeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	cake := initCake(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"title": "Lemon Melon",
		"description": "A cheesecake with melon",
		"rating": 4,
		"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
	}`)
	request := httptest.NewRequest(http.MethodPut, baseUrl+"/cakes/"+strconv.Itoa(cake.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "ROOT")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	responseBody := transformResponse(response)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	data := responseBody["data"].(map[string]interface{})
	assert.Equal(t, cake.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Lemon Melon", data["title"])
	assert.Equal(t, "A cheesecake with melon", data["description"])
	assert.Equal(t, 4.0, data["rating"])
	assert.Equal(t, "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", data["image"])
}

func TestUpdateCakeFailed(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	cake := initCake(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"title":""}`)
	request := httptest.NewRequest(http.MethodPut, baseUrl+"/cakes/"+strconv.Itoa(cake.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "ROOT")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	responseBody := transformResponse(response)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad request", responseBody["status"])
}

func TestGetCakeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	cake := initCake(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, baseUrl+"/cakes/"+strconv.Itoa(cake.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "ROOT")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	responseBody := transformResponse(response)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	data := responseBody["data"].(map[string]interface{})
	assert.Equal(t, "Lemon cheesecake", data["title"])
	assert.Equal(t, "A cheesecake made of lemon", data["description"])
	assert.Equal(t, 7.3, data["rating"])
	assert.Equal(t, "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", data["image"])
}

func TestGetCakeFailed(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, baseUrl+"/cakes/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "ROOT")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	responseBody := transformResponse(response)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Data not found", responseBody["status"])
}

func TestDeleteCakeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	cake := initCake(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, baseUrl+"/cakes/"+strconv.Itoa(cake.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "ROOT")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	responseBody := transformResponse(response)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCakeFailed(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, baseUrl+"/cakes/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "ROOT")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	responseBody := transformResponse(response)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Data not found", responseBody["status"])
}

func TestListCakeSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	cake1 := initCake(db)
	cake2 := initCake(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, baseUrl+"/cakes", nil)
	request.Header.Add("Content-Type", "applicGeMethodGetation/json")
	request.Header.Add("X-API-Key", "ROOT")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	responseBody := transformResponse(response)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var cakes = responseBody["data"].([]interface{})
	cakeResponse1 := cakes[0].(map[string]interface{})
	cakeResponse2 := cakes[1].(map[string]interface{})

	assert.Equal(t, cake1.Id, int(cakeResponse1["id"].(float64)))
	assert.Equal(t, cake1.Title, cakeResponse1["title"])

	assert.Equal(t, cake2.Id, int(cakeResponse2["id"].(float64)))
	assert.Equal(t, cake2.Title, cakeResponse2["title"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCake(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, baseUrl+"/cakes", nil)
	request.Header.Add("Content-Type", "applicGeMethodGetation/json")
	request.Header.Add("X-API-Key", "")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	responseBody := transformResponse(response)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
}
