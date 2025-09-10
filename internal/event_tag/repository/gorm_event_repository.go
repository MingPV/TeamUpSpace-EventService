package repository

import (
	"github.com/MingPV/EventService/internal/entities"
	"gorm.io/gorm"
)

type GormEventTagRepository struct {
	db *gorm.DB
}

func NewGormEventTagRepository(db *gorm.DB) EventTagRepository {
	return &GormEventTagRepository{db: db}
}

func (r *GormEventTagRepository) Save(event_tag *entities.EventTag) error {
	return r.db.Create(&event_tag).Error
}

func (r *GormEventTagRepository) FindAll() ([]*entities.EventTag, error) {
	var eventTagValues []entities.EventTag
	if err := r.db.Find(&eventTagValues).Error; err != nil {
		return nil, err
	}

	event_tags := make([]*entities.EventTag, len(eventTagValues))
	for i := range eventTagValues {
		event_tags[i] = &eventTagValues[i]
	}
	return event_tags, nil
}

func (r *GormEventTagRepository) FindByEventAndTagID(event_id int, tag_id int) (*entities.EventTag, error) {
	var event_tag entities.EventTag
	if err := r.db.First(&event_tag, "event_id = ? AND tag_id = ?", event_id, tag_id).Error; err != nil {
		return &entities.EventTag{}, err
	}
	return &event_tag, nil
}

func (r *GormEventTagRepository) FindByEventID(event_id int) ([]*entities.EventTag, error) {
	var eventTagValues []entities.EventTag
	if err := r.db.Where("event_id = ?", event_id).Find(&eventTagValues).Error; err != nil {
		return nil, err
	}

	event_tags := make([]*entities.EventTag, len(eventTagValues))
	for i := range eventTagValues {
		event_tags[i] = &eventTagValues[i]
	}
	return event_tags, nil
}

func (r *GormEventTagRepository) FindByTagID(tag_id int) ([]*entities.EventTag, error) {
	var eventTagValues []entities.EventTag
	if err := r.db.Where("tag_id = ?", tag_id).Find(&eventTagValues).Error; err != nil {
		return nil, err
	}

	event_tags := make([]*entities.EventTag, len(eventTagValues))
	for i := range eventTagValues {
		event_tags[i] = &eventTagValues[i]
	}
	return event_tags, nil
}

func (r *GormEventTagRepository) Delete(event_id int, tag_id int) error {
	result := r.db.Delete(&entities.EventTag{}, "event_id = ? AND tag_id = ?", event_id, tag_id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
