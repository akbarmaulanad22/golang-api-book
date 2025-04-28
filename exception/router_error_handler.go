package exception

import (
	"api_book/helper"
	"api_book/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func RouterErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if errValidation, ok := err.(validator.ValidationErrors); ok {
		badRequestHandler(w, r, errValidation.Error())
	} else if errNotFound, ok := err.(NotFoundError); ok {
		notFoundHandler(w, r, errNotFound.Error)
	} else {
		internalServerHandler(w, r, err)
	}
}

func internalServerHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Data: err,
	}

	helper.WriteToResponse(w, webResponse)
}

func badRequestHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	webResponse := web.WebResponse{
		Code: http.StatusBadRequest,
		Message: "BAD REQUEST",
		Data: err,
	}

	helper.WriteToResponse(w, webResponse)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	webResponse := web.WebResponse{
		Code: http.StatusNotFound,
		Message: "NOT FOUND",
		Data: err,
	}

	helper.WriteToResponse(w, webResponse)
}