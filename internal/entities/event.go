package entities

import "time"

type Event struct {
	ID              	uint      	`gorm:"primaryKey;autoIncrement" json:"id"`
	EventName       	string    	`json:"event_name"`
	EventDescription 	string    	`json:"event_description"`
	StartAt         	string    	`json:"start_at"`
	EndAt           	string    	`json:"end_at"`
	MainImageURL    	string    	`json:"main_image_url"`
	RegisterStartDt 	string    	`json:"register_start_dt"`
	RegisterCloseDt 	string    	`json:"register_close_dt"`
	CreatedAt      		time.Time	`json:"created_at"`
	UpdatedAt      		time.Time	`gorm:"autoUpdateTime" json:"updated_at"`
}

type EventTag struct {
	EventID     uint   		`json:"event_id"`
	TagID       uint   		`json:"tag_id"`
	CreatedAt   time.Time	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time	`gorm:"autoUpdateTime" json:"updated_at"`
}

type Tag struct {
	ID          uint      	`gorm:"primaryKey;autoIncrement" json:"id"`
	TagName     string    	`json:"tag_name"`
	CreatedAt   time.Time	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time	`gorm:"autoUpdateTime" json:"updated_at"`
}
