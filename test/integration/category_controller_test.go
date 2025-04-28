package integration

import (
	"api_book/controller"
	"api_book/exception"
	"api_book/helper"
	"api_book/middleware"
	"api_book/model/domain"
	"api_book/repository"
	"api_book/service"
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
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func setupDBTesting() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_api_book_testing")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	
	return db
}

func setupRouter(db *sql.DB) http.Handler {
	baseUrl := "/api/v1"
	
	validator := validator.New()
	router := httprouter.New()

	// categories routes
	
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validator)
	categoryController := controller.NewCategoryController(categoryService)

	router.GET(baseUrl + "/categories", categoryController.FindAll)
	router.POST(baseUrl + "/categories", categoryController.Create)

	router.GET(baseUrl + "/categories/:categoryId", categoryController.FindById)
	router.PUT(baseUrl + "/categories/:categoryId", categoryController.Update)
	router.DELETE(baseUrl + "/categories/:categoryId", categoryController.Delete)

	// books routes

	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validator)
	bookController := controller.NewBookController(bookService)

	router.GET(baseUrl + "/books", bookController.FindAll)
	router.POST(baseUrl + "/books", bookController.Create)

	router.GET(baseUrl + "/books/:bookId", bookController.FindById)
	router.PUT(baseUrl + "/books/:bookId", bookController.Update)
	router.DELETE(baseUrl + "/books/:bookId", bookController.Delete)

	router.PanicHandler = exception.RouterErrorHandler

	return middleware.NewAuthMiddleware(router)
}

func truncateCategories(db *sql.DB)  {
	db.Exec("truncate categories")
}

func TestCategoryCreateSuccess(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)
	
	requestBody := strings.NewReader(`{
		"name": "Self taught",
		"slug": "self-taught",
		"is_active": true
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:9090/api/v1/categories", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestCategoryCreateFailed(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)
	
	requestBody := strings.NewReader(`{}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:9090/api/v1/categories", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestCategoryUpdateSuccess(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)

	tx, _ := db.Begin()

	repo := repository.NewCategoryRepository()
	category := repo.Save(context.Background(), tx, domain.Category{
		Name: "rabdom",
		Slug: "rdao",
		IsActive: false,
	})
	
	tx.Commit()
	
	requestBody := strings.NewReader(`{
		"name": "Self taught",
		"slug": "self-taught",
		"is_active": true
	}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:9090/api/v1/categories/" + strconv.Itoa(category.Id) , requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestCategoryUpdateErrorBadRequest(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)

	tx, _ := db.Begin()

	repo := repository.NewCategoryRepository()
	category := repo.Save(context.Background(), tx, domain.Category{
		Name: "rabdom",
		Slug: "rdao",
		IsActive: false,
	})
	
	tx.Commit()
	
	requestBody := strings.NewReader(`{}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:9090/api/v1/categories/" + strconv.Itoa(category.Id) , requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestCategoryUpdateErrorNotFound(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)

	requestBody := strings.NewReader(`{
		"name": "Self taught",
		"slug": "self-taught",
		"is_active": true
	}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:9090/api/v1/categories/404", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
}

func TestCategoryDeleteSuccess(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)

	tx, _ := db.Begin()

	repo := repository.NewCategoryRepository()
	category := repo.Save(context.Background(), tx, domain.Category{
		Name: "rabdom",
		Slug: "rdao",
		IsActive: false,
	})
	
	tx.Commit()
	

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:9090/api/v1/categories/" + strconv.Itoa(category.Id), nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestCategoryDeleteFailed(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:9090/api/v1/categories/404", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
}

func TestCategoryFindByIdSucces(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)

	tx, _ := db.Begin()

	repo := repository.NewCategoryRepository()
	category := repo.Save(context.Background(), tx, domain.Category{
		Name: "rabdom",
		Slug: "rdao",
		IsActive: false,
	})
	
	tx.Commit()
	

	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/api/v1/categories/" + strconv.Itoa(category.Id), nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestCategoryFindByIdFailed(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/api/v1/categories/404", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
}

func TestCategoryFindAll(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)

	tx, _ := db.Begin()

	repo := repository.NewCategoryRepository()
	category1 := repo.Save(context.Background(), tx, domain.Category{
		Name: "rabdom",
		Slug: "rdao",
		IsActive: false,
	})

	category2 := repo.Save(context.Background(), tx, domain.Category{
		Name: "rabdom",
		Slug: "rdao",
		IsActive: false,
	})
	
	tx.Commit()
	

	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/api/v1/categories", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	bodyResponse := map[string]interface{}{}

	bodyByte, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(bodyByte, &bodyResponse)

	
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, category1.Name, bodyResponse["data"].([]interface{})[0].(map[string]interface{})["name"])
	assert.Equal(t, category2.Name, bodyResponse["data"].([]interface{})[1].(map[string]interface{})["name"])
}

func TestCategoryApiKey(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateCategories(db)	

	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/api/v1/categories", nil)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusUnauthorized, recorder.Result().StatusCode)
}