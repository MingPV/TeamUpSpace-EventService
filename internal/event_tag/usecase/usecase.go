package usecase

import (
	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/event_tag/repository"
)

// EventTagService
type EventTagService struct {
	repo repository.EventTagRepository
}

// Init EventTagService function
func NewEventTagService(repo repository.EventTagRepository) EventTagUseCase {
	return &EventTagService{repo: repo}
}

// EventTagService Methods - 1 create
func (s *EventTagService) CreateEventTag(event_tag *entities.EventTag) error {
	if err := s.repo.Save(event_tag); err != nil {
		return err
	}
	return nil
}

// EventTagService Methods - 2 find all
func (s *EventTagService) FindAllEventTags() ([]*entities.EventTag, error) {
	event_tags, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return event_tags, nil
}

// EventTagService Methods - 3 find by id
func (s *EventTagService) FindEventTagByID(event_id int, tag_id int) (*entities.EventTag, error) {
	return s.repo.FindByID(event_id, tag_id)
}

// EventTagService Methods - 5 delete
func (s *EventTagService) DeleteEventTag(event_id int, tag_id int) error {
	if err := s.repo.Delete(event_id, tag_id); err != nil {
		return err
	}
	return nil
}
