package usecase

import "github.com/MingPV/EventService/internal/entities"

type EventUseCase interface {
	FindAllEvents() ([]*entities.Event, error)
	CreateEvent(event *entities.Event) error
	PatchEvent(id int, event *entities.Event) (*entities.Event, error)
	DeleteEvent(id int) error
	FindEventByID(id int) (*entities.Event, error)
}