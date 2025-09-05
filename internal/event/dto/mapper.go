package dto

import "github.com/MingPV/EventService/internal/entities"

func ToEventResponse(event *entities.Event) *EventResponse {
	return &EventResponse{
		ID:               event.ID,
		EventName:        event.EventName,
		EventDescription: event.EventDescription,
		StartAt:          event.StartAt,
		EndAt:            event.EndAt,
		MainImageUrl:     event.MainImageUrl,
		RegisterStartDt:  event.RegisterStartDt,
		RegisterCloseDt:  event.RegisterCloseDt,
		CreatedAt:        event.CreatedAt,
		UpdatedAt:        event.UpdatedAt,
	}
}

func ToEventResponseList(events []*entities.Event) []*EventResponse {
	result := make([]*EventResponse, 0, len(events))
	for _, e := range events {
		result = append(result, ToEventResponse(e))
	}
	return result
}
