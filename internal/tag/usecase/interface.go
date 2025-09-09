package usecase

import "github.com/MingPV/EventService/internal/entities"

type TagUseCase interface {
	FindAllTags() ([]*entities.Tag, error)
	CreateTag(tag *entities.Tag) error
	PatchTag(id int, tag *entities.Tag) (*entities.Tag, error)
	DeleteTag(id int) error
	FindTagByID(id int) (*entities.Tag, error)
}