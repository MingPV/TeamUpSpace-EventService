package rest

import (
	"strconv"

	// "github.com/MingPV/EventService/pkg/apperror"

	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/event/dto"
	"github.com/MingPV/EventService/internal/event/usecase"
	responses "github.com/MingPV/EventService/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpEventHandler struct {
	eventUseCase usecase.EventUseCase
}

func NewHttpEventHandler(useCase usecase.EventUseCase) *HttpEventHandler {
	return &HttpEventHandler{eventUseCase: useCase}
}

// CreateEvent godoc
// @Summary Create a new event
// @Tags events
// @Accept json
// @Produce json
// @Param event body dto.CreateEventRequest true "Event payload"
// @Success 201 {object} dto.EventResponse
// @Router /events [post]
func (h *HttpEventHandler) CreateEvent(c *fiber.Ctx) error {
	var req dto.CreateEventRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	event := &entities.Event{
		EventName:        req.EventName,
		EventDescription: req.EventDescription,
		StartAt:          req.StartAt,
		EndAt:            req.EndAt,
		MainImageURL:     req.MainImageURL,
		RegisterStartDt:  req.RegisterStartDt,
		RegisterCloseDt:  req.RegisterCloseDt,
	}
	if err := h.eventUseCase.CreateEvent(event); err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToEventResponse(event))
}

// FindAllEvents godoc
// @Summary Get all events
// @Tags events
// @Produce json
// @Success 200 {array} entities.Event
// @Router /events [get]
func (h *HttpEventHandler) FindAllEvents(c *fiber.Ctx) error {
	events, err := h.eventUseCase.FindAllEvents()
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToEventResponseList(events))
}

// FindEventByID godoc
// @Summary Get event by ID
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} entities.Event
// @Router /events/{id} [get]
func (h *HttpEventHandler) FindEventByID(c *fiber.Ctx) error {
	id := c.Params("id")
	eventID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	event, err := h.eventUseCase.FindEventByID(eventID)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToEventResponse(event))
}

// PatchEvent godoc
// @Summary Update an event partially
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param event body entities.Event true "Event update payload"
// @Success 200 {object} entities.Event
// @Router /events/{id} [patch]
func (h *HttpEventHandler) PatchEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	eventID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	var req dto.CreateEventRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	event := &entities.Event{
		EventName:        req.EventName,
		EventDescription: req.EventDescription,
		StartAt:          req.StartAt,
		EndAt:            req.EndAt,
		MainImageURL:     req.MainImageURL,
		RegisterStartDt:  req.RegisterStartDt,
		RegisterCloseDt:  req.RegisterCloseDt,
	}

	msg, err := validatePatchEvent(event)
	if err != nil {
		return responses.ErrorWithMessage(c, err, msg)
	}

	updatedEvent, err := h.eventUseCase.PatchEvent(eventID, event)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToEventResponse(updatedEvent))
}

// DeleteEvent godoc
// @Summary Delete an event by ID
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} responses.MessageResponse
// @Router /events/{id} [delete]
func (h *HttpEventHandler) DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	eventID, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	if err := h.eventUseCase.DeleteEvent(eventID); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "event deleted")
}

func validatePatchEvent(event *entities.Event) (string, error) {

	// if event.EventName == "" {
	// 	return "event name is required", apperror.ErrInvalidData
	// }
	// if event.StartAt == "" {
	// 	return "start at is required", apperror.ErrInvalidData
	// }
	// if event.EndAt == "" {
	// 	return "end at is required", apperror.ErrInvalidData
	// }
	// if event.RegisterStartDt == "" {
	// 	return "register start date is required", apperror.ErrInvalidData
	// }
	// if event.RegisterCloseDt == "" {
	// 	return "register close date is required", apperror.ErrInvalidData
	// }

	return "", nil
}
