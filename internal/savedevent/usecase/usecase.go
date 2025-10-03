package usecase

import (
	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/savedevent/repository"
	"github.com/google/uuid"
)

type SavedEventService struct {
	repo repository.SavedEventRepository
}

func NewSavedEventService(repo repository.SavedEventRepository) SavedEventUseCase {
	return &SavedEventService{repo: repo}
}

func (s *SavedEventService) SaveEvent(event *entities.SavedEvent) (*entities.SavedEvent, error) {
	if err := s.repo.Save(event); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *SavedEventService) FindSavedEvent(userID string, eventID int32) (*entities.SavedEvent, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByUserAndEvent(uid, int(eventID))
}

func (s *SavedEventService) FindAllByUser(userID string) ([]*entities.SavedEvent, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAllByUser(uid)
}

func (s *SavedEventService) FindAllByEvent(eventID int32) ([]*entities.SavedEvent, error) {
	return s.repo.FindAllByEvent(int(eventID))
}

func (s *SavedEventService) DeleteSavedEvent(userID string, eventID int32) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	return s.repo.Delete(uid, int(eventID))
}
