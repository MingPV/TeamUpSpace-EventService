package grpc

import (
	"context"

	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/savedevent/usecase"
	"github.com/MingPV/EventService/pkg/apperror"
	savedeventpb "github.com/MingPV/EventService/proto/savedevent"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
)

type GrpcSavedEventHandler struct {
	savedEventUseCase usecase.SavedEventUseCase
	savedeventpb.UnimplementedSavedEventServiceServer
}

func NewGrpcSavedEventHandler(uc usecase.SavedEventUseCase) *GrpcSavedEventHandler {
	return &GrpcSavedEventHandler{savedEventUseCase: uc}
}

func (h *GrpcSavedEventHandler) SaveEvent(ctx context.Context, req *savedeventpb.SaveEventRequest) (*savedeventpb.SaveEventResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid user_id UUID")
	}

	se := &entities.SavedEvent{
		UserID:   userID,
		EventID:  int(req.EventId),
	}

	created, err := h.savedEventUseCase.SaveEvent(se)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &savedeventpb.SaveEventResponse{
		SavedEvent: toProtoSavedEvent(created),
	}, nil
}

func (h *GrpcSavedEventHandler) FindSavedEvent(ctx context.Context, req *savedeventpb.FindSavedEventRequest) (*savedeventpb.FindSavedEventResponse, error) {
	se, err := h.savedEventUseCase.FindSavedEvent(req.UserId, req.EventId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &savedeventpb.FindSavedEventResponse{
		SavedEvent: toProtoSavedEvent(se),
	}, nil
}

func (h *GrpcSavedEventHandler) FindAllByUser(ctx context.Context, req *savedeventpb.FindAllByUserRequest) (*savedeventpb.FindAllByUserResponse, error) {
	events, err := h.savedEventUseCase.FindAllByUser(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoEvents []*savedeventpb.SavedEvent
	for _, e := range events {
		protoEvents = append(protoEvents, toProtoSavedEvent(e))
	}

	return &savedeventpb.FindAllByUserResponse{
		SavedEvents: protoEvents,
	}, nil
}

func (h *GrpcSavedEventHandler) FindAllByEvent(ctx context.Context, req *savedeventpb.FindAllByEventRequest) (*savedeventpb.FindAllByEventResponse, error) {
	events, err := h.savedEventUseCase.FindAllByEvent(req.EventId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoEvents []*savedeventpb.SavedEvent
	for _, e := range events {
		protoEvents = append(protoEvents, toProtoSavedEvent(e))
	}

	return &savedeventpb.FindAllByEventResponse{
		SavedEvents: protoEvents,
	}, nil
}

func (h *GrpcSavedEventHandler) DeleteSavedEvent(ctx context.Context, req *savedeventpb.DeleteSavedEventRequest) (*savedeventpb.DeleteSavedEventResponse, error) {
	if err := h.savedEventUseCase.DeleteSavedEvent(req.UserId, req.EventId); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &savedeventpb.DeleteSavedEventResponse{
		Message: "saved event deleted",
	}, nil
}

func toProtoSavedEvent(se *entities.SavedEvent) *savedeventpb.SavedEvent {
	return &savedeventpb.SavedEvent{
		UserId:    se.UserID.String(),
		EventId:   int32(se.EventID),
		CreatedAt: se.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
