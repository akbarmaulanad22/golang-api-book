package helper

import (
	"api_book/model/domain"
	"api_book/model/web"
)

func ConvertToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
		Slug: category.Slug,
		IsActive: category.IsActive,
	}
}

func ConvertToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	categoriesResponses := []web.CategoryResponse{}
	
	for _, category := range categories {
		categoriesResponses = append(categoriesResponses, ConvertToCategoryResponse(category))
	}

	return categoriesResponses
}

func ConvertToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		Id: book.Id,
		Title: book.Title,
		Author: book.Author,
		Publisher: book.Publisher,
		Year: book.Year,
		Genre: book.Genre,
		Pages: book.Pages,
		// CategoryId: book.CategoryId,
	}
}

func ConvertToBookResponses(books []domain.Book) []web.BookResponse {
	booksResponses := []web.BookResponse{}
	
	for _, book := range books {
		booksResponses = append(booksResponses, ConvertToBookResponse(book))
	}

	return booksResponses
}