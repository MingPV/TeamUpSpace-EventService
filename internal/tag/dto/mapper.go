package dto

import "github.com/MingPV/EventService/internal/entities"

func ToTagResponse(tag *entities.Tag) *TagResponse {
	return &TagResponse{
		ID:        tag.ID,
		TagName:   tag.TagName,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
	}
}

func ToTagResponseList(tags []*entities.Tag) []*TagResponse {
	result := make([]*TagResponse, 0, len(tags))
	for _, e := range tags {
		result = append(result, ToTagResponse(e))
	}
	return result
}
