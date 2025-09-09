package entities

import "time"

type Event struct {
	ID              	uint      	`gorm:"primaryKey;autoIncrement" json:"id"`
	EventName       	string    	`json:"event_name"`
	EventDescription 	string    	`json:"event_description"`
	StartAt         	string    	`json:"start_at"`
	EndAt           	string    	`json:"end_at"`
	MainImageUrl    	string    	`json:"main_image_url"`
	RegisterStartDt 	string    	`json:"register_start_dt"`
	RegisterCloseDt 	string    	`json:"register_close_dt"`
	CreatedAt      		time.Time	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      		time.Time	`gorm:"autoUpdateTime" json:"updated_at"`

	Tags []Tag `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags,omitempty"`
}

type EventTag struct {
	EventID     uint   		`gorm:"primaryKey" json:"event_id"`
	TagID       uint   		`gorm:"primaryKey" json:"tag_id"`
	CreatedAt   time.Time	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time	`gorm:"autoUpdateTime" json:"updated_at"`

	// Foreign key relationships
	Event Event `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"event"`
	Tag   Tag   `gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tag"`
}

type Tag struct {
	ID          uint      	`gorm:"primaryKey;autoIncrement" json:"id"`
	TagName     string    	`json:"tag_name"`
	CreatedAt   time.Time	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time	`gorm:"autoUpdateTime" json:"updated_at"`

	Events []Event `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"events,omitempty"`
}
