package controller

import (
	"api_book/helper"
	"api_book/model/web"
	"api_book/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type bookControllerImplement struct {
	Service service.BookService
}

func NewBookController(service service.BookService) BookController {
	return &bookControllerImplement{
		Service: service,
	}
}


// Create implements BookController.
func (controller *bookControllerImplement) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookRequest := web.BookCreateRequest{}
	helper.ReadFromRequest(r, &bookRequest)

	bookResponse := controller.Service.Create(r.Context(), bookRequest)	

	response := web.WebResponse{
		Code: http.StatusOK,
		Message: "OK",
		Data: bookResponse,
	}

	helper.WriteToResponse(w, response)
	
}

// Delete implements BookController.
func (controller *bookControllerImplement) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	paramId := p.ByName("bookId")
	id, errIdNotValid := strconv.Atoi(paramId)
	helper.PanicIfError(errIdNotValid)

	controller.Service.Delete(r.Context(), id)
	
	response := web.WebResponse{
		Code: http.StatusOK,
		Message: "OK",
		Data: nil,
	}

	helper.WriteToResponse(w, response)
}

// FindAll implements BookController.
func (controller *bookControllerImplement) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	booksResponses := controller.Service.FindAll(r.Context())

	response := web.WebResponse{
		Code: http.StatusOK,
		Message: "OK",
		Data: booksResponses,
	}

	helper.WriteToResponse(w, response)
	
}

// FindById implements BookController.
func (controller *bookControllerImplement) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	paramId := p.ByName("bookId")
	id, errIdNotValid := strconv.Atoi(paramId)
	helper.PanicIfError(errIdNotValid)

	bookResponse := controller.Service.FindById(r.Context(), id)

	response := web.WebResponse{
		Code: http.StatusOK,
		Message: "OK",
		Data: bookResponse,
	}

	helper.WriteToResponse(w, response)	
	
}

// Update implements BookController.
func (controller *bookControllerImplement) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookRequest := web.BookUpdateRequest{}
	helper.ReadFromRequest(r, &bookRequest)
	
	paramId := p.ByName("bookId")
	id, errIdNotValid := strconv.Atoi(paramId)
	helper.PanicIfError(errIdNotValid)

	bookRequest.Id = int(id)
	bookResponse := controller.Service.Update(r.Context(), bookRequest)

	response := web.WebResponse{
		Code: http.StatusOK,
		Message: "OK",
		Data: bookResponse,
	}

	helper.WriteToResponse(w, response)	
}

