package service

import (
	"api_book/model/web"
	"context"
)

type BookService interface {
	Create(ctx context.Context, bookRequest web.BookCreateRequest) web.BookResponse
	Update(ctx context.Context, bookRequest web.BookUpdateRequest) web.BookResponse
	Delete(ctx context.Context, categoryId int) 
	FindById(ctx context.Context, categoryId int) web.BookResponse
	FindAll(ctx context.Context) []web.BookResponse
}