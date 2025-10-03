package usecase

import "github.com/MingPV/EventService/internal/entities"

type SavedEventUseCase interface {
	SaveEvent(event *entities.SavedEvent) (*entities.SavedEvent, error)
	FindSavedEvent(userID string, eventID int32) (*entities.SavedEvent, error)
	FindAllByUser(userID string) ([]*entities.SavedEvent, error)
	FindAllByEvent(eventID int32) ([]*entities.SavedEvent, error)
	DeleteSavedEvent(userID string, eventID int32) error
}
