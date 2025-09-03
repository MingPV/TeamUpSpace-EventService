package repository

import (
	"github.com/MingPV/EventService/internal/entities"
	"gorm.io/gorm"
)

type GormEventRepository struct {
	db *gorm.DB
}

func NewGormEventRepository(db *gorm.DB) EventRepository {
	return &GormEventRepository{db: db}
}

func (r *GormEventRepository) Save(event *entities.Event) error {
	return r.db.Create(&event).Error
}

func (r *GormEventRepository) FindAll() ([]*entities.Event, error) {
	var eventValues []entities.Event
	if err := r.db.Find(&eventValues).Error; err != nil {
		return nil, err
	}

	events := make([]*entities.Event, len(eventValues))
	for i := range eventValues {
		events[i] = &eventValues[i]
	}
	return events, nil
}

func (r *GormEventRepository) FindByID(id int) (*entities.Event, error) {
	var event entities.Event
	if err := r.db.First(&event, id).Error; err != nil {
		return &entities.Event{}, err
	}
	return &event, nil
}

func (r *GormEventRepository) Patch(id int, event *entities.Event) error {
	result := r.db.Model(&entities.Event{}).Where("id = ?", id).Updates(event)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormEventRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Event{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
