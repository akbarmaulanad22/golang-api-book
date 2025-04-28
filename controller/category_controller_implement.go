package controller

import (
	"api_book/helper"
	"api_book/model/web"
	"api_book/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type categoryControllerImplement struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryControllerImplement{
		CategoryService: categoryService,
	}
}

// Create implements CategoryController.
func (controller *categoryControllerImplement) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequest(r, &categoryRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryRequest)

	response := web.WebResponse{
		Code: http.StatusOK,
		Message: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponse(w, response)
}

// Delete implements CategoryController.
func (controller *categoryControllerImplement) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	paramId := p.ByName("categoryId")
	id, errIdNotValid := strconv.Atoi(paramId)
	helper.PanicIfError(errIdNotValid)

	controller.CategoryService.Delete(r.Context(), id)
	response := web.WebResponse{
		Code: http.StatusOK,
		Message: "OK",
		Data: nil,
	}

	helper.WriteToResponse(w, response)
}

// FindAll implements CategoryController.
func (controller *categoryControllerImplement) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoriesResponses := controller.CategoryService.FindAll(r.Context())

	response := web.WebResponse{
		Code: http.StatusOK,
		Message: "OK",
		Data: categoriesResponses,
	}
	
	helper.WriteToResponse(w, response)

}

// FindById implements CategoryController.
func (controller *categoryControllerImplement) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	paramId := p.ByName("categoryId")
	id, errIdNotValid := strconv.Atoi(paramId)
	helper.PanicIfError(errIdNotValid)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	response := web.WebResponse{
		Code: http.StatusOK,
		Message: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponse(w, response)
}

// Update implements CategoryController.
func (controller *categoryControllerImplement) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequest(r, &categoryRequest)
	
	paramId := p.ByName("categoryId")
	id, errIdNotValid := strconv.Atoi(paramId)
	helper.PanicIfError(errIdNotValid)

	categoryRequest.Id = id
	categoryResponse := controller.CategoryService.Update(r.Context(), categoryRequest)
	
	response := web.WebResponse{
		Code: http.StatusOK,
		Message: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponse(w, response)
}


