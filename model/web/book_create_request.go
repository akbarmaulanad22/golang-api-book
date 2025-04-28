package web

type BookCreateRequest struct {
	Title      string `json:"title" validate:"required"`
	Author     string `json:"author" validate:"required"`
	Publisher  string `json:"publisher" validate:"required"`
	Year       int    `json:"year" validate:"required"`
	Genre      string `json:"genre" validate:"required"`
	Pages      int    `json:"pages" validate:"required"`
	CategoryId int    `json:"category_id" validate:"required"`
}