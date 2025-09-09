package dto

import "time"

type EventTagResponse struct {
	EventID   uint      `json:"event_id"`
	TagID     uint      `json:"tag_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
