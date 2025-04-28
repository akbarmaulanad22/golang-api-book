package app

import (
	"api_book/controller"
	"api_book/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController, bookController controller.BookController) *httprouter.Router {
	baseUrl := "/api/v1"
	
	router := httprouter.New()

	// categories routes

	router.GET(baseUrl + "/categories", categoryController.FindAll)
	router.POST(baseUrl + "/categories", categoryController.Create)

	router.GET(baseUrl + "/categories/:categoryId", categoryController.FindById)
	router.PUT(baseUrl + "/categories/:categoryId", categoryController.Update)
	router.DELETE(baseUrl + "/categories/:categoryId", categoryController.Delete)

	// books routes

	router.GET(baseUrl + "/books", bookController.FindAll)
	router.POST(baseUrl + "/books", bookController.Create)

	router.GET(baseUrl + "/books/:bookId", bookController.FindById)
	router.PUT(baseUrl + "/books/:bookId", bookController.Update)
	router.DELETE(baseUrl + "/books/:bookId", bookController.Delete)

	router.PanicHandler = exception.RouterErrorHandler

	return router
}