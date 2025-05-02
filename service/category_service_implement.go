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

type categoryServiceImplement struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
	Logger *logrus.Logger
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate, logger *logrus.Logger) CategoryService {
	return &categoryServiceImplement{
		CategoryRepository: categoryRepository, 
		DB: db, 
		Validate: validate,
		Logger: logger,
	}
}


// Create implements CategoryService.
func (service *categoryServiceImplement) Create(ctx context.Context, categoryRequest web.CategoryCreateRequest) web.CategoryResponse {
	errValidation := service.Validate.Struct(categoryRequest)
	helper.PanicIfError(errValidation)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	
	category := service.CategoryRepository.Save(ctx, tx, domain.Category{
		Name: categoryRequest.Name,
		Slug: categoryRequest.Slug,
		IsActive: categoryRequest.IsActive,
	})

	service.Logger.WithField("data", category).Info("category created successfully")
	
	return helper.ConvertToCategoryResponse(category)
}

// Delete implements CategoryService.
func (service *categoryServiceImplement) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	
	category, errNotFound := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if errNotFound != nil {
		panic(exception.NewNotFoundError(errNotFound))
	}

	service.CategoryRepository.Delete(ctx, tx, categoryId)

	service.Logger.WithField("data", category).Info("category deleted successfully")
	
}

// FindAll implements CategoryService.
func (service *categoryServiceImplement) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ConvertToCategoryResponses(categories)
}

// FindById implements CategoryService.
func (service *categoryServiceImplement) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, errNotFound := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if errNotFound != nil {
		panic(exception.NewNotFoundError(errNotFound))
	}

	return helper.ConvertToCategoryResponse(category)
}

// Update implements CategoryService.
func (service *categoryServiceImplement) Update(ctx context.Context, categoryRequest web.CategoryUpdateRequest) web.CategoryResponse {
	
	errValidation := service.Validate.Struct(categoryRequest)
	helper.PanicIfError(errValidation)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, errNotFound := service.CategoryRepository.FindById(ctx, tx, categoryRequest.Id)
	if errNotFound != nil {
		panic(exception.NewNotFoundError(errNotFound))
	}

	category.Id = categoryRequest.Id
	category.Name = categoryRequest.Name
	category.Slug = categoryRequest.Slug
	category.IsActive = categoryRequest.IsActive

	category = service.CategoryRepository.Update(ctx, tx, category)

	service.Logger.WithField("data", category).Info("category updated successfully")

	return helper.ConvertToCategoryResponse(category)
}

