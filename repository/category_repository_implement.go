package repository

import (
	"api_book/helper"
	"api_book/model/domain"
	"context"
	"database/sql"
	"errors"
)

type categoryRepositoryImplement struct {
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepositoryImplement{}
}

func (repo *categoryRepositoryImplement) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	result, err := tx.ExecContext(ctx, "INSERT INTO categories (name, slug, is_active) VALUES (?, ?, ?)", category.Name, category.Slug, category.IsActive)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repo *categoryRepositoryImplement) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	_, err := tx.ExecContext(ctx, "UPDATE categories SET name = ?, slug = ?, is_active = ? WHERE id = ?", category.Name, category.Slug, category.IsActive, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repo *categoryRepositoryImplement) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	_, err := tx.ExecContext(ctx, "DELETE FROM categories WHERE id = ?", categoryId)
	helper.PanicIfError(err)
}

func (repo *categoryRepositoryImplement) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	row, err := tx.QueryContext(ctx, "SELECT id, name, slug, is_active FROM categories WHERE id = ? LIMIT 1", categoryId)
	helper.PanicIfError(err)
	defer row.Close()

	category := domain.Category{}

	if !row.Next() {
		return category, errors.New("category not found")
	} 
	
	errScan := row.Scan(&category.Id, &category.Name, &category.Slug, &category.IsActive)
	helper.PanicIfError(errScan)
	
	return category, nil
}

func (repo *categoryRepositoryImplement) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	rows, err := tx.QueryContext(ctx, "SELECT id, name, slug, is_active FROM categories")
	helper.PanicIfError(err)
	defer rows.Close()

	categories := []domain.Category{}

	for rows.Next() {
		category := domain.Category{}

		errScan := rows.Scan(&category.Id, &category.Name, &category.Slug, &category.IsActive)
		helper.PanicIfError(errScan)

		categories = append(categories, category)
	} 
	
	return categories
}