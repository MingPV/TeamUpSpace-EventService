package usecase

import (
	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/tag/repository"
)

// TagService Interface
type TagService struct {
	repo repository.TagRepository
}

// Init TagService function
func NewTagService(repo repository.TagRepository) TagUseCase {
	return &TagService{repo: repo}
}

// TagService Methods - 1 create
func (s *TagService) CreateTag(tag *entities.Tag) error {
	if err := s.repo.Save(tag); err != nil {
		return err
	}
	return nil
}

// TagService Methods - 2 find all
func (s *TagService) FindAllTags() ([]*entities.Tag, error) {
	tags, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// TagService Methods - 3 find by id
func (s *TagService) FindTagByID(id int) (*entities.Tag, error) {

	tag, err := s.repo.FindByID(id)
	if err != nil {
		return &entities.Tag{}, err
	}
	return tag, nil
}

// TagService Methods - 4 patch
func (s *TagService) PatchTag(id int, tag *entities.Tag) (*entities.Tag, error) {

	if err := s.repo.Patch(id, tag); err != nil {
		return nil, err
	}
	updatedTag, _ := s.repo.FindByID(id)

	return updatedTag, nil
}

// TagService Methods - 5 delete
func (s *TagService) DeleteTag(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
