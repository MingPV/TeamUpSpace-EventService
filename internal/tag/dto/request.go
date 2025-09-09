package dto

type CreateTagRequest struct {
	TagName   string    `json:"tag_name" validate:"required"`
}
