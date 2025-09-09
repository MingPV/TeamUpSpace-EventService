package dto

import "github.com/MingPV/EventService/internal/entities"

func ToEventTagResponse(event *entities.EventTag) *EventTagResponse {
	return &EventTagResponse{
		EventID: event.EventID,
		TagID:   event.TagID,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}
}

func ToEventTagResponseList(events []*entities.EventTag) []*EventTagResponse {
	result := make([]*EventTagResponse, 0, len(events))
	for _, e := range events {
		result = append(result, ToEventTagResponse(e))
	}
	return result
}
