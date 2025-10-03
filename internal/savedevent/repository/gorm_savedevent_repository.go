package repository

import (
	"github.com/MingPV/EventService/internal/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormSavedEventRepository struct {
	db *gorm.DB
}

func NewGormSavedEventRepository(db *gorm.DB) SavedEventRepository {
	return &GormSavedEventRepository{db: db}
}

func (r *GormSavedEventRepository) Save(event *entities.SavedEvent) error {
	return r.db.Create(event).Error
}

func (r *GormSavedEventRepository) FindByUserAndEvent(userID uuid.UUID, eventID int) (*entities.SavedEvent, error) {
	var se entities.SavedEvent
	if err := r.db.Where("user_id = ? AND event_id = ?", userID, eventID).First(&se).Error; err != nil {
		return nil, err
	}
	return &se, nil
}

func (r *GormSavedEventRepository) FindAllByUser(userID uuid.UUID) ([]*entities.SavedEvent, error) {
	var values []entities.SavedEvent
	if err := r.db.Where("user_id = ?", userID).Find(&values).Error; err != nil {
		return nil, err
	}
	events := make([]*entities.SavedEvent, len(values))
	for i := range values {
		events[i] = &values[i]
	}
	return events, nil
}

func (r *GormSavedEventRepository) FindAllByEvent(eventID int) ([]*entities.SavedEvent, error) {
	var values []entities.SavedEvent
	if err := r.db.Where("event_id = ?", eventID).Find(&values).Error; err != nil {
		return nil, err
	}
	events := make([]*entities.SavedEvent, len(values))
	for i := range values {
		events[i] = &values[i]
	}
	return events, nil
}

func (r *GormSavedEventRepository) Delete(userID uuid.UUID, eventID int) error {
	result := r.db.Where("user_id = ? AND event_id = ?", userID, eventID).Delete(&entities.SavedEvent{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
