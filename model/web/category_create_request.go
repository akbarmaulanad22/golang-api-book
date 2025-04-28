package web

type CategoryCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Slug     string `json:"slug" validate:"required"`
	IsActive bool   `json:"is_active" validate:"boolean"`
}