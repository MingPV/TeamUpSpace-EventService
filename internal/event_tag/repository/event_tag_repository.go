package repository

import "github.com/MingPV/EventService/internal/entities"

type EventTagRepository interface {
	Save(event_tag *entities.EventTag) error
	FindAll() ([]*entities.EventTag, error)
	FindByID(event_id int, tag_id int) (*entities.EventTag, error)
	Delete(event_id int, tag_id int) error
}
