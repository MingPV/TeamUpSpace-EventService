package dto

type EventResponse struct {
	ID               uint	`json:"id"`
	EventName        string	`json:"event_name"`
	EventDescription string	`json:"event_description"`
	StartAt          string `json:"start_at"`
	EndAt            string `json:"end_at"`
	MainImageURL     string	`json:"main_image_url"`
	RegisterStartDt  string `json:"register_start_dt"`
	RegisterCloseDt  string `json:"register_close_dt"`
}
