package repository

import (
	"github.com/MingPV/EventService/internal/entities"
	"github.com/google/uuid"
)

type SavedEventRepository interface {
	Save(event *entities.SavedEvent) error
	FindByUserAndEvent(userID uuid.UUID, eventID int) (*entities.SavedEvent, error)
	FindAllByUser(userID uuid.UUID) ([]*entities.SavedEvent, error)
	FindAllByEvent(eventID int) ([]*entities.SavedEvent, error)
	Delete(userID uuid.UUID, eventID int) error
}
