package usecase

import "github.com/MingPV/EventService/internal/entities"

type EventTagUseCase interface {
	FindAllEventTags() ([]*entities.EventTag, error)
	CreateEventTag(event_tag *entities.EventTag) error
	FindByEventAndTagID(event_id int, tag_id int) (*entities.EventTag, error)
	FindByEventID(event_id int) ([]*entities.EventTag, error)
	FindByTagID(tag_id int) ([]*entities.EventTag, error)
	DeleteEventTag(event_id int, tag_id int) error
}