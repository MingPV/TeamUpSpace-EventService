package grpc

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/event/event.proto

import (
	"context"

	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/event/usecase"
	"github.com/MingPV/EventService/pkg/apperror"
	eventpb "github.com/MingPV/EventService/proto/event"
	"google.golang.org/grpc/status"
)

type GrpcEventHandler struct {
	eventUseCase usecase.EventUseCase
	eventpb.UnimplementedEventServiceServer
}

func NewGrpcEventHandler(uc usecase.EventUseCase) *GrpcEventHandler {
	return &GrpcEventHandler{eventUseCase: uc}
}

func (h *GrpcEventHandler) CreateEvent(ctx context.Context, req *eventpb.CreateEventRequest) (*eventpb.CreateEventResponse, error) {
	event := &entities.Event{
		EventName:  req.EventName,
		EventDescription: req.EventDescription,
		StartAt:    req.StartAt,
		EndAt:      req.EndAt,
		MainImageUrl: req.MainImageUrl,
		RegisterStartDt: req.RegisterStartDt,
		RegisterCloseDt: req.RegisterCloseDt,
	}
	if err := h.eventUseCase.CreateEvent(event); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &eventpb.CreateEventResponse{Event: toProtoEvent(event)}, nil
}

func (h *GrpcEventHandler) FindEventByID(ctx context.Context, req *eventpb.FindEventByIDRequest) (*eventpb.FindEventByIDResponse, error) {
	event, err := h.eventUseCase.FindEventByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &eventpb.FindEventByIDResponse{Event: toProtoEvent(event)}, nil
}

func (h *GrpcEventHandler) FindAllEvents(ctx context.Context, req *eventpb.FindAllEventsRequest) (*eventpb.FindAllEventsResponse, error) {
	events, err := h.eventUseCase.FindAllEvents()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoEvents []*eventpb.Event
	for _, o := range events {
		protoEvents = append(protoEvents, toProtoEvent(o))
	}

	return &eventpb.FindAllEventsResponse{Events: protoEvents}, nil
}

func (h *GrpcEventHandler) PatchEvent(ctx context.Context, req *eventpb.PatchEventRequest) (*eventpb.PatchEventResponse, error) {
	event := &entities.Event{
		EventName:  req.EventName,
		EventDescription: req.EventDescription,
		StartAt:    req.StartAt,
		EndAt:      req.EndAt,
		MainImageUrl: req.MainImageUrl,
		RegisterStartDt: req.RegisterStartDt,
		RegisterCloseDt: req.RegisterCloseDt,
	}
	updatedEvent, err := h.eventUseCase.PatchEvent(int(req.Id), event)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &eventpb.PatchEventResponse{Event: toProtoEvent(updatedEvent)}, nil
}

func (h *GrpcEventHandler) DeleteEvent(ctx context.Context, req *eventpb.DeleteEventRequest) (*eventpb.DeleteEventResponse, error) {
	if err := h.eventUseCase.DeleteEvent(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &eventpb.DeleteEventResponse{Message: "event deleted"}, nil
}

// helper function convert entities.Event to eventpb.Event
func toProtoEvent(e *entities.Event) *eventpb.Event {
	return &eventpb.Event{
		Id:               uint32(e.ID),
		EventName:        e.EventName,
		EventDescription: e.EventDescription,
		StartAt:          e.StartAt,
		EndAt:            e.EndAt,
		MainImageUrl:     e.MainImageUrl,
		RegisterStartDt:  e.RegisterStartDt,
		RegisterCloseDt:  e.RegisterCloseDt,
		CreatedAt:        e.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        e.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
