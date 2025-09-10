package rest

import (
	"strconv"

	// "github.com/MingPV/EventService/pkg/apperror"

	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/event_tag/dto"
	"github.com/MingPV/EventService/internal/event_tag/usecase"
	responses "github.com/MingPV/EventService/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpEventTagHandler struct {
	eventUseCase usecase.EventTagUseCase
}

func NewHttpEventTagHandler(useCase usecase.EventTagUseCase) *HttpEventTagHandler {
	return &HttpEventTagHandler{eventUseCase: useCase}
}

// CreateEventTag godoc
// @Summary Create a new event tag
// @Tags event_tags
// @Accept json
// @Produce json
// @Param event body dto.CreateEventTagRequest true "Event Tag payload"
// @Success 201 {object} dto.EventTagResponse
// @Router /event_tags [post]
func (h *HttpEventTagHandler) CreateEventTag(c *fiber.Ctx) error {
	var req dto.CreateEventTagRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	event := &entities.EventTag{
		EventID: req.EventID,
		TagID:   req.TagID,
	}
	if err := h.eventUseCase.CreateEventTag(event); err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToEventTagResponse(event))
}

// FindAllEventTags godoc
// @Summary Get all event tags
// @Tags event_tags
// @Produce json
// @Success 200 {array} entities.EventTag
// @Router /event_tags [get]
func (h *HttpEventTagHandler) FindAllEventTags(c *fiber.Ctx) error {
	eventTags, err := h.eventUseCase.FindAllEventTags()
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToEventTagResponseList(eventTags))
}

// FindByEventAndTagID godoc
// @Summary Get event tag by ID
// @Tags event_tags
// @Produce json
// @Param event_id path int true "Event ID"
// @Param tag_id path int true "Tag ID"
// @Success 200 {object} entities.EventTag
// @Router /event_tags/{event_id}/{tag_id} [get]
func (h *HttpEventTagHandler) FindByEventAndTagID(c *fiber.Ctx) error {
	eventID := c.Params("event_id")
	tagID := c.Params("tag_id")

	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid event id")
	}

	tagIDInt, err := strconv.Atoi(tagID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid tag id")
	}

	eventTag, err := h.eventUseCase.FindByEventAndTagID(eventIDInt, tagIDInt)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToEventTagResponse(eventTag))
}

// FindByEventID godoc
// @Summary Get event tags by Event ID
// @Tags event_tags
// @Produce json
// @Param event_id path int true "Event ID"
// @Success 200 {array} entities.EventTag
// @Router /event_tags/event/{event_id} [get]
func (h *HttpEventTagHandler) FindByEventID(c *fiber.Ctx) error {
	eventID := c.Params("event_id")

	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid event id")
	}

	eventTags, err := h.eventUseCase.FindByEventID(eventIDInt)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToEventTagResponseList(eventTags))
}

// FindByTagID godoc
// @Summary Get event tags by Tag ID
// @Tags event_tags
// @Produce json
// @Param tag_id path int true "Tag ID"
// @Success 200 {array} entities.EventTag
// @Router /event_tags/tag/{tag_id} [get]
func (h *HttpEventTagHandler) FindByTagID(c *fiber.Ctx) error {
	tagID := c.Params("tag_id")

	tagIDInt, err := strconv.Atoi(tagID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid tag id")
	}

	eventTags, err := h.eventUseCase.FindByTagID(tagIDInt)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToEventTagResponseList(eventTags))
}

// DeleteEventTag godoc
// @Summary Delete an event tag by ID
// @Tags event_tags
// @Produce json
// @Param event_id path int true "Event ID"
// @Param tag_id path int true "Tag ID"
// @Success 200 {object} responses.MessageResponse
// @Router /event_tags/{event_id}/{tag_id} [delete]
func (h *HttpEventTagHandler) DeleteEventTag(c *fiber.Ctx) error {
	eventID := c.Params("event_id")
	tagID := c.Params("tag_id")

	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid event id")
	}

	tagIDInt, err := strconv.Atoi(tagID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid tag id")
	}

	if err := h.eventUseCase.DeleteEventTag(eventIDInt, tagIDInt); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "event tag deleted")
}
