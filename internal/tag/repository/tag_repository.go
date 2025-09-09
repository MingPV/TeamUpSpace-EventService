package repository

import "github.com/MingPV/EventService/internal/entities"

type TagRepository interface {
	Save(tag *entities.Tag) error
	FindAll() ([]*entities.Tag, error)
	FindByID(id int) (*entities.Tag, error)
	Patch(id int, tag *entities.Tag) error
	Delete(id int) error
}
