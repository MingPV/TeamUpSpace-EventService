package repository

import "github.com/MingPV/EventService/internal/entities"

type EventRepository interface {
	Save(event *entities.Event) error
	FindAll() ([]*entities.Event, error)
	FindByID(id int) (*entities.Event, error)
	Patch(id int, event *entities.Event) error
	Delete(id int) error
}
