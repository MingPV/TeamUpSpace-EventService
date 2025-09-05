package dto

type CreateEventRequest struct {
	EventName        string    `json:"event_name" validate:"required"`
	EventDescription string    `json:"event_description" validate:"required"`
	StartAt          string    `json:"start_at" validate:"required"`
	EndAt            string    `json:"end_at" validate:"required"`
	MainImageUrl     string    `json:"main_image_url" validate:"required,url"`
	RegisterStartDt  string    `json:"register_start_dt" validate:"required"`
	RegisterCloseDt  string    `json:"register_close_dt" validate:"required"`
}
