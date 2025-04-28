package web

type CategoryUpdateRequest struct {
	Id       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Slug     string `json:"slug" validate:"required"`
	IsActive bool   `json:"is_active" validate:"boolean"`
}