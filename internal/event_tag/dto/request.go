package dto

type CreateEventTagRequest struct {
	EventID uint `json:"event_id" validate:"required"`
	TagID   uint `json:"tag_id" validate:"required"`
}
