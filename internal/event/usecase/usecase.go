package usecase

import (
	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/event/repository"
)

// EventService
type EventService struct {
	repo repository.EventRepository
}

// Init EventService function
func NewEventService(repo repository.EventRepository) EventUseCase {
	return &EventService{repo: repo}
}

// EventService Methods - 1 create
func (s *EventService) CreateEvent(event *entities.Event) error {
	if err := s.repo.Save(event); err != nil {
		return err
	}
	return nil
}

// EventService Methods - 2 find all
func (s *EventService) FindAllEvents() ([]*entities.Event, error) {
	events, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return events, nil
}

// EventService Methods - 3 find by id
func (s *EventService) FindEventByID(id int) (*entities.Event, error) {

	event, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.Event{}, err
	}
	return event, nil
}

// EventService Methods - 4 patch
func (s *EventService) PatchEvent(id int, event *entities.Event) (*entities.Event, error) {

	if err := s.repo.Patch(id, event); err != nil {
		return nil, err
	}
	updatedEvent, _ := s.repo.FindByID(id)

	return updatedEvent, nil
}

// EventService Methods - 5 delete
func (s *EventService) DeleteEvent(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
