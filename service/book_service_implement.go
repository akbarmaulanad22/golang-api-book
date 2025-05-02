package service

import (
	"api_book/exception"
	"api_book/helper"
	"api_book/model/domain"
	"api_book/model/web"
	"api_book/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type bookServiceImplement struct {
	Repository 	repository.BookRepository
	DB         	*sql.DB
	Validate   	*validator.Validate
	Logger 		*logrus.Logger
}

func NewBookService(repo repository.BookRepository, db *sql.DB, validate *validator.Validate, logger *logrus.Logger) BookService {
	return &bookServiceImplement{
		Repository: repo,
		DB:         db,
		Validate:   validate,
		Logger: 	logger,
	}
}

// Create implements BookService.
func (service *bookServiceImplement) Create(ctx context.Context, bookRequest web.BookCreateRequest) web.BookResponse {
	errValidation := service.Validate.Struct(bookRequest)
	helper.PanicIfError(errValidation)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book := service.Repository.Save(ctx, tx, domain.Book{
		Title: bookRequest.Title,
		Author: bookRequest.Author,
		Publisher: bookRequest.Publisher,
		Year: bookRequest.Year,
		Genre: bookRequest.Genre,
		Pages: bookRequest.Pages,
		CategoryId: bookRequest.CategoryId,
	})

	service.Logger.WithField("data", book).Info("book created successfully")
	
	return helper.ConvertToBookResponse(book)
}

// Delete implements BookService.
func (service *bookServiceImplement) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, errNotFound := service.Repository.FindById(ctx, tx, categoryId)
	if errNotFound != nil {
		panic(exception.NewNotFoundError(errNotFound))
	}

	service.Repository.Delete(ctx, tx, categoryId)
	
	service.Logger.WithField("data", book).Info("book deleted successfully")
}

// FindAll implements BookService.
func (service *bookServiceImplement) FindAll(ctx context.Context) []web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	books := service.Repository.FindAll(ctx, tx)

	return helper.ConvertToBookResponses(books)
}

// FindById implements BookService.
func (service *bookServiceImplement) FindById(ctx context.Context, categoryId int) web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	book, errNotFound := service.Repository.FindById(ctx, tx, categoryId)
	if errNotFound != nil {
		panic(exception.NewNotFoundError(errNotFound))
	}

	return helper.ConvertToBookResponse(book)
}

// Update implements BookService.
func (service *bookServiceImplement) Update(ctx context.Context, bookRequest web.BookUpdateRequest) web.BookResponse {
	errValidation := service.Validate.Struct(bookRequest)
	helper.PanicIfError(errValidation)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, errNotFound := service.Repository.FindById(ctx, tx, bookRequest.Id)
	if errNotFound != nil {
		panic(exception.NewNotFoundError(errNotFound))
	}

	book.Title = bookRequest.Title
	book.Author = bookRequest.Author
	book.Publisher = bookRequest.Publisher
	book.Year = bookRequest.Year
	book.Genre = bookRequest.Genre
	book.Pages = bookRequest.Pages
	book.CategoryId = bookRequest.CategoryId

	book = service.Repository.Update(ctx, tx, book)

	service.Logger.WithField("data", book).Info("book updated successfully")
	
	return helper.ConvertToBookResponse(book)
}


