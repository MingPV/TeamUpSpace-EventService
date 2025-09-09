package rest

import (
	"strconv"

	// "github.com/MingPV/EventService/pkg/apperror"

	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/tag/dto"
	"github.com/MingPV/EventService/internal/tag/usecase"
	responses "github.com/MingPV/EventService/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpTagHandler struct {
	tagUseCase usecase.TagUseCase
}

func NewHttpTagHandler(useCase usecase.TagUseCase) *HttpTagHandler {
	return &HttpTagHandler{tagUseCase: useCase}
}

// CreateTag godoc
// @Summary Create a new tag
// @Tags tags
// @Accept json
// @Produce json
// @Param tag body dto.CreateTagRequest true "Tag payload"
// @Success 201 {object} dto.TagResponse
// @Router /tags [post]
func (h *HttpTagHandler) CreateTag(c *fiber.Ctx) error {
	var req dto.CreateTagRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	tag := &entities.Tag{
		TagName: req.TagName,
	}
	if err := h.tagUseCase.CreateTag(tag); err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToTagResponse(tag))
}

// FindAllTags godoc
// @Summary Get all tags
// @Tags tags
// @Produce json
// @Success 200 {array} entities.Tag
// @Router /tags [get]
func (h *HttpTagHandler) FindAllTags(c *fiber.Ctx) error {
	tags, err := h.tagUseCase.FindAllTags()
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToTagResponseList(tags))
}

// FindTagByID godoc
// @Summary Get tag by ID
// @Tags tags
// @Produce json
// @Param id path int true "Tag ID"
// @Success 200 {object} entities.Tag
// @Router /tags/{id} [get]
func (h *HttpTagHandler) FindTagByID(c *fiber.Ctx) error {
	id := c.Params("id")
	tagID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	tag, err := h.tagUseCase.FindTagByID(tagID)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToTagResponse(tag))
}

// PatchTag godoc
// @Summary Update a tag partially
// @Tags tags
// @Accept json
// @Produce json
// @Param id path int true "Tag ID"
// @Param tag body dto.CreateTagRequest true "Tag update payload"
// @Success 200 {object} entities.Tag
// @Router /tags/{id} [patch]
func (h *HttpTagHandler) PatchTag(c *fiber.Ctx) error {
	id := c.Params("id")
	tagID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	var req dto.CreateTagRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	tag := &entities.Tag{
		TagName: req.TagName,
	}

	// msg, err := validatePatchTag(tag)
	// if err != nil {
	// 	return responses.ErrorWithMessage(c, err, msg)
	// }

	updatedTag, err := h.tagUseCase.PatchTag(tagID, tag)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToTagResponse(updatedTag))
}

// DeleteTag godoc
// @Summary Delete a tag by ID
// @Tags tags
// @Produce json
// @Param id path int true "Tag ID"
// @Success 200 {object} responses.MessageResponse
// @Router /tags/{id} [delete]
func (h *HttpTagHandler) DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")
	tagID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	if err := h.tagUseCase.DeleteTag(tagID); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "tag deleted")
}

// func validatePatchTag(tag *entities.Tag) (string, error) {

// 	// if tag.TagName == "" {
// 	// 	return "tag name is required", apperror.ErrInvalidData
// 	// }

// 	return "", nil
// }
