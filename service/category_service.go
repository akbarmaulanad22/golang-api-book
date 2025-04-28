package service

import (
	"api_book/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, categoryRequest web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, categoryRequest web.CategoryUpdateRequest) web.CategoryResponse
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
}