package entities

import (
	"time"

	"github.com/google/uuid"
)

type SavedEvent struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	EventID   int       `gorm:"primaryKey" json:"event_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Event *Event `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE" json:"event"`
}
