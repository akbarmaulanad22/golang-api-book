package integration

import (
	"api_book/model/domain"
	"api_book/repository"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


func truncateBooks(db *sql.DB)  {
	db.Exec("truncate books")
}

func TestBookCreateSuccess(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateBooks(db)
	
	requestBody := strings.NewReader(`
		{
			"title": "no love just lift edit ahaahah",
			"author": "barr edit",
			"publisher": "pio edit",
			"year": 2025,
			"genre": "gatau edit ahahaha",
			"pages": 2,
			"category_id": 2
		}
	`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:9090/api/v1/books", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestBookCreateFailed(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateBooks(db)
	
	requestBody := strings.NewReader(`{}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:9090/api/v1/books", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestBookUpdateSuccess(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateBooks(db)

	tx, _ := db.Begin()

	repo := repository.NewBookRepository()
	category := repo.Save(context.Background(), tx, domain.Book{
		Title: "no love just lift ed",
		Author: "barr",
		Publisher: "pio",
		Year: 2025,
		Genre: "gatau",
		Pages: 2,
		CategoryId: 2,
	})
	
	tx.Commit()
	
	requestBody := strings.NewReader(`{
		"title": "no love just lift edit",
		"author": "barr edit",
		"publisher": "pio edit",
		"year": 2025,
		"genre": "gatau edit",
		"pages": 2,
		"category_id": 2
	}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:9090/api/v1/books/" + strconv.Itoa(category.Id) , requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestBookUpdateFailedBadRequest(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateBooks(db)

	tx, _ := db.Begin()

	repo := repository.NewBookRepository()
	category := repo.Save(context.Background(), tx, domain.Book{
		Title: "no love just lift",
		Author: "barr",
		Publisher: "pio",
		Year: 2025,
		Genre: "gatau",
		Pages: 2,
		CategoryId: 2,
	})
	
	tx.Commit()
	
	requestBody := strings.NewReader(`{}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:9090/api/v1/books/" + strconv.Itoa(category.Id) , requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestBookUpdateFailedNotFound(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateBooks(db)
	
	requestBody := strings.NewReader(`{
		"title": "no love just lift edit",
		"author": "barr edit",
		"publisher": "pio edit",
		"year": 2025,
		"genre": "gatau edit",
		"pages": 2,
		"category_id": 2
	}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:9090/api/v1/books/404", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
}

func TestBookDeleteSuccess(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateBooks(db)

	tx, _ := db.Begin()

	repo := repository.NewBookRepository()
	category := repo.Save(context.Background(), tx, domain.Book{
		Title: "no love just lift",
		Author: "barr",
		Publisher: "pio",
		Year: 2025,
		Genre: "gatau",
		Pages: 2,
		CategoryId: 2,
	})
	
	tx.Commit()
	
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:9090/api/v1/books/" + strconv.Itoa(category.Id) , nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestBookDeleteFailed(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateBooks(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:9090/api/v1/categories/404", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
}

func TestBookFindByIdSuccess(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateBooks(db)

	tx, _ := db.Begin()

	repo := repository.NewBookRepository()
	book := repo.Save(context.Background(), tx, domain.Book{
		Title: "no love",
		Author: "barr",
		Publisher: "pio",
		Year: 2025,
		Genre: "gatau",
		Pages: 2,
		CategoryId: 2,
	})
	
	tx.Commit()
	
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/api/v1/books/" + strconv.Itoa(book.Id), nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestBookFindByIdFailed(t *testing.T) {

	db := setupDBTesting()
	router := setupRouter(db)
	truncateBooks(db)
	
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/api/v1/books/404", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
}

func TestFindAllBooks(t *testing.T) {
	db := setupDBTesting()
	router := setupRouter(db)
	truncateBooks(db)

	tx, _ := db.Begin()

	repo := repository.NewBookRepository()
	book1 := repo.Save(context.Background(), tx, domain.Book{
		Title: "no love 1",
		Author: "barr",
		Publisher: "pio",
		Year: 2025,
		Genre: "gatau",
		Pages: 2,
		CategoryId: 2,
	})

	book2 := repo.Save(context.Background(), tx, domain.Book{
		Title: "no love 2",
		Author: "barr",
		Publisher: "pio",
		Year: 2025,
		Genre: "gatau",
		Pages: 2,
		CategoryId: 2,
	})
	
	tx.Commit()
	

	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/api/v1/books", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "KEPOBEJIR")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	bodyResponse := map[string]interface{}{}

	bodyByte, _ := io.ReadAll(recorder.Result().Body)
	json.Unmarshal(bodyByte, &bodyResponse)

	
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, book1.Title, bodyResponse["data"].([]interface{})[0].(map[string]interface{})["title"])
	assert.Equal(t, book2.Title, bodyResponse["data"].([]interface{})[1].(map[string]interface{})["title"])
}