package repository

import (
	"github.com/MingPV/EventService/internal/entities"
	"gorm.io/gorm"
)

type GormTagRepository struct {
	db *gorm.DB
}

func NewGormTagRepository(db *gorm.DB) TagRepository {
	return &GormTagRepository{db: db}
}

func (r *GormTagRepository) Save(tag *entities.Tag) error {
	return r.db.Create(&tag).Error
}

func (r *GormTagRepository) FindAll() ([]*entities.Tag, error) {
	var tagValues []entities.Tag
	if err := r.db.Find(&tagValues).Error; err != nil {
		return nil, err
	}

	tags := make([]*entities.Tag, len(tagValues))
	for i := range tagValues {
		tags[i] = &tagValues[i]
	}
	return tags, nil
}

func (r *GormTagRepository) FindByID(id int) (*entities.Tag, error) {
	var tag entities.Tag
	if err := r.db.First(&tag, id).Error; err != nil {
		return &entities.Tag{}, err
	}
	return &tag, nil
}

func (r *GormTagRepository) Patch(id int, tag *entities.Tag) error {
	result := r.db.Model(&entities.Tag{}).Where("id = ?", id).Updates(tag)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormTagRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Tag{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
