package grpc

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/event/event.proto

import (
	"context"

	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/event_tag/usecase"
	"github.com/MingPV/EventService/pkg/apperror"
	event_tag_pb "github.com/MingPV/EventService/proto/event_tag"
	"google.golang.org/grpc/status"
)

type GrpcEventTagHandler struct {
	eventTagUseCase usecase.EventTagUseCase
	event_tag_pb.UnimplementedEventTagServiceServer
}

func NewGrpcEventTagHandler(uc usecase.EventTagUseCase) *GrpcEventTagHandler {
	return &GrpcEventTagHandler{eventTagUseCase: uc}
}

func (h *GrpcEventTagHandler) CreateEventTag(ctx context.Context, req *event_tag_pb.CreateEventTagRequest) (*event_tag_pb.CreateEventTagResponse, error) {
	event := &entities.EventTag{
		EventID:   uint(req.EventId),
		TagID:     uint(req.TagId),
	}
	if err := h.eventTagUseCase.CreateEventTag(event); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &event_tag_pb.CreateEventTagResponse{EventTag: toProtoEventTag(event)}, nil
}
func (h *GrpcEventTagHandler) FindEventTagByID(ctx context.Context, req *event_tag_pb.FindEventTagByIDRequest) (*event_tag_pb.FindEventTagByIDResponse, error) {
	event, err := h.eventTagUseCase.FindEventTagByID(int(req.EventId), int(req.TagId))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &event_tag_pb.FindEventTagByIDResponse{EventTag: toProtoEventTag(event)}, nil
}

func (h *GrpcEventTagHandler) FindAllEventTags(ctx context.Context, req *event_tag_pb.FindAllEventTagsRequest) (*event_tag_pb.FindAllEventTagsResponse, error) {
	events, err := h.eventTagUseCase.FindAllEventTags()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoEvents []*event_tag_pb.EventTag
	for _, o := range events {
		protoEvents = append(protoEvents, toProtoEventTag(o))
	}

	return &event_tag_pb.FindAllEventTagsResponse{EventTags: protoEvents}, nil
}

func (h *GrpcEventTagHandler) DeleteEventTag(ctx context.Context, req *event_tag_pb.DeleteEventTagRequest) (*event_tag_pb.DeleteEventTagResponse, error) {
	if err := h.eventTagUseCase.DeleteEventTag(int(req.EventId), int(req.TagId)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &event_tag_pb.DeleteEventTagResponse{Message: "event tag deleted"}, nil
}

// helper function convert entities.Event to eventpb.Event
func toProtoEventTag(et *entities.EventTag) *event_tag_pb.EventTag {
	return &event_tag_pb.EventTag{
		EventId:   uint32(et.EventID),
		TagId:     uint32(et.TagID),
		CreatedAt: et.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: et.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}