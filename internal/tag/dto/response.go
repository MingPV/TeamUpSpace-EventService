package dto

import "time"

type TagResponse struct {
	ID               uint		`json:"id"`
	TagName          string		`json:"tag_name"`
	CreatedAt        time.Time 	`json:"created_at"`
	UpdatedAt        time.Time 	`json:"updated_at"`
}
