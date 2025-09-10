package repository

import "github.com/MingPV/EventService/internal/entities"

type EventTagRepository interface {
	Save(event_tag *entities.EventTag) error
	FindAll() ([]*entities.EventTag, error)
	FindByEventAndTagID(event_id int, tag_id int) (*entities.EventTag, error)
	FindByEventID(event_id int) ([]*entities.EventTag, error)
	FindByTagID(tag_id int) ([]*entities.EventTag, error)
	Delete(event_id int, tag_id int) error
}
