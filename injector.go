//go:build wireinject
// +build wireinject

package main

import (
	"api_book/app"
	"api_book/controller"
	"api_book/middleware"
	"api_book/repository"
	"api_book/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

var bookSet = wire.NewSet(
	repository.NewBookRepository,
	service.NewBookService,
	controller.NewBookController,
)

func InitializedServer() *http.Server {
	

	wire.Build(
		app.NewDB, 
		app.NewValidator,
		app.NewLogger,
		categorySet,
		bookSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		app.NewServer,
	)
	
	return nil
}