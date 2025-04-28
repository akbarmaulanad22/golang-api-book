package repository

import (
	"api_book/helper"
	"api_book/model/domain"
	"context"
	"database/sql"
	"errors"
)

type bookRepositoryImplement struct {
}

func NewBookRepository() BookRepository {
	return &bookRepositoryImplement{}
}

// Delete implements BookRepository.
func (repo *bookRepositoryImplement) Delete(ctx context.Context, tx *sql.Tx, bookId int) {	
	_, err := tx.ExecContext(ctx, "DELETE FROM books WHERE id = ?", bookId)
	helper.PanicIfError(err)
}

// FindAll implements BookRepository.
func (repo *bookRepositoryImplement) FindAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	rows, err := tx.QueryContext(ctx, "SELECT id, title, author, publisher, year, genre, pages, category_id FROM books")
	helper.PanicIfError(err)
	defer rows.Close()

	books := []domain.Book{}
	for rows.Next() {
		book := domain.Book{}
		errScan := rows.Scan(
			&book.Id, 
			&book.Title, 
			&book.Author, 
			&book.Publisher, 
			&book.Year, 
			&book.Genre, 
			&book.Pages, 
			&book.CategoryId,
		)

		helper.PanicIfError(errScan)

		books = append(books, book)
	}

	return books
}

// FindById implements BookRepository.
func (repo *bookRepositoryImplement) FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error) {
	row, err := tx.QueryContext(ctx, "SELECT id, title, author, publisher, year, genre, pages, category_id FROM books WHERE id = ? LIMIT 1", bookId)
	helper.PanicIfError(err)
	defer row.Close()
	
	book := domain.Book{}
	if !row.Next() {
		return book, errors.New("book not found")
	}

	errScan := row.Scan(
		&book.Id, 
		&book.Title, 
		&book.Author, 
		&book.Publisher, 
		&book.Year, 
		&book.Genre, 
		&book.Pages, 
		&book.CategoryId,
	)

	helper.PanicIfError(errScan)

	return book, nil
}

// Save implements BookRepository.
func (repo *bookRepositoryImplement) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	result, err := tx.ExecContext(
		ctx, 
		"INSERT INTO books (title, author, publisher, year, genre, pages, category_id) VALUES (?, ?, ?, ?, ?, ?, ?)", 
		book.Title, book.Author, book.Publisher, book.Year, book.Genre, book.Pages, book.CategoryId,
	)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	book.Id = int(id)
	return book
}

// Update implements BookRepository.
func (repo *bookRepositoryImplement) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	_, err := tx.ExecContext(
		ctx, 
		"UPDATE books SET title = ?, author = ?, publisher = ?, year = ?, genre = ?, pages = ?, category_id = ?", 
		book.Title, book.Author, book.Publisher, book.Year, book.Genre, book.Pages, book.CategoryId,
	)

	helper.PanicIfError(err)

	return book
}

