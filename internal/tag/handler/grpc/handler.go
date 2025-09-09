package grpc

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/event/event.proto

import (
	"context"

	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/tag/usecase"
	"github.com/MingPV/EventService/pkg/apperror"
	tagpb "github.com/MingPV/EventService/proto/tag"
	"google.golang.org/grpc/status"
)

type GrpcTagHandler struct {
	tagUseCase usecase.TagUseCase
	tagpb.UnimplementedTagServiceServer
}

func NewGrpcTagHandler(uc usecase.TagUseCase) *GrpcTagHandler {
	return &GrpcTagHandler{tagUseCase: uc}
}

func (h *GrpcTagHandler) CreateTag(ctx context.Context, req *tagpb.CreateTagRequest) (*tagpb.CreateTagResponse, error) {
	tag := &entities.Tag{
		TagName: req.TagName,
	}
	if err := h.tagUseCase.CreateTag(tag); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &tagpb.CreateTagResponse{Tag: toProtoTag(tag)}, nil
}

func (h *GrpcTagHandler) FindTagByID(ctx context.Context, req *tagpb.FindTagByIDRequest) (*tagpb.FindTagByIDResponse, error) {
	tag, err := h.tagUseCase.FindTagByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &tagpb.FindTagByIDResponse{Tag: toProtoTag(tag)}, nil
}

func (h *GrpcTagHandler) FindAllTags(ctx context.Context, req *tagpb.FindAllTagsRequest) (*tagpb.FindAllTagsResponse, error) {
	tags, err := h.tagUseCase.FindAllTags()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoTags []*tagpb.Tag
	for _, o := range tags {
		protoTags = append(protoTags, toProtoTag(o))
	}

	return &tagpb.FindAllTagsResponse{Tags: protoTags}, nil
}

func (h *GrpcTagHandler) PatchTag(ctx context.Context, req *tagpb.PatchTagRequest) (*tagpb.PatchTagResponse, error) {
	tag := &entities.Tag{
		TagName: req.TagName,
	}
	updatedTag, err := h.tagUseCase.PatchTag(int(req.Id), tag)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &tagpb.PatchTagResponse{Tag: toProtoTag(updatedTag)}, nil
}
func (h *GrpcTagHandler) DeleteTag(ctx context.Context, req *tagpb.DeleteTagRequest) (*tagpb.DeleteTagResponse, error) {
	if err := h.tagUseCase.DeleteTag(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &tagpb.DeleteTagResponse{Message: "tag deleted"}, nil
}

// helper function convert entities.Tag to tagpb.Tag
func toProtoTag(t *entities.Tag) *tagpb.Tag {
	return &tagpb.Tag{
		Id:        uint32(t.ID),
		TagName:   t.TagName,
		CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: t.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}