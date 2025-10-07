package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// // Order
	// orderHandler "github.com/MingPV/EventService/internal/order/handler/rest"
	// orderRepository "github.com/MingPV/EventService/internal/order/repository"
	// orderUseCase "github.com/MingPV/EventService/internal/order/usecase"

	// Event
	eventHandler "github.com/MingPV/EventService/internal/event/handler/rest"
	eventRepository "github.com/MingPV/EventService/internal/event/repository"
	eventUseCase "github.com/MingPV/EventService/internal/event/usecase"
	"github.com/MingPV/EventService/pkg/mq"

	// Tag
	tagHandler "github.com/MingPV/EventService/internal/tag/handler/rest"
	tagRepository "github.com/MingPV/EventService/internal/tag/repository"
	tagUseCase "github.com/MingPV/EventService/internal/tag/usecase"

	// EventTag
	eventTagHandler "github.com/MingPV/EventService/internal/event_tag/handler/rest"
	eventTagRepository "github.com/MingPV/EventService/internal/event_tag/repository"
	eventTagUseCase "github.com/MingPV/EventService/internal/event_tag/usecase"
)

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB) {

	api := app.Group("/api/v1")

	// === Dependency Wiring ===

	// // Order
	// orderRepo := orderRepository.NewGormOrderRepository(db)
	// orderService := orderUseCase.NewOrderService(orderRepo)
	// orderHandler := orderHandler.NewHttpOrderHandler(orderService)

	// MQ Publisher
	// rabbitURL := "amqp://guest:guest@host.docker.internal:5672/"
	rabbitURL := "amqp://guest:guest@localhost:5672/"
	mqPublisher := mq.NewRabbitMQPublisher(rabbitURL)

	// Event
	eventRepo := eventRepository.NewGormEventRepository(db)
	eventService := eventUseCase.NewEventService(eventRepo, mqPublisher)
	eventHandler := eventHandler.NewHttpEventHandler(eventService)

	// Tag
	tagRepo := tagRepository.NewGormTagRepository(db)
	tagService := tagUseCase.NewTagService(tagRepo)
	tagHandler := tagHandler.NewHttpTagHandler(tagService)

	// EventTag
	eventTagRepo := eventTagRepository.NewGormEventTagRepository(db)
	eventTagService := eventTagUseCase.NewEventTagService(eventTagRepo)
	eventTagHandler := eventTagHandler.NewHttpEventTagHandler(eventTagService)

	// === Public Routes ===

	// // Order routes
	// orderGroup := api.Group("/orders")
	// orderGroup.Get("/", orderHandler.FindAllOrders)
	// orderGroup.Get("/:id", orderHandler.FindOrderByID)
	// orderGroup.Post("/", orderHandler.CreateOrder)
	// orderGroup.Patch("/:id", orderHandler.PatchOrder)
	// orderGroup.Delete("/:id", orderHandler.DeleteOrder)

	// Event routes
	eventGroup := api.Group("/events")
	eventGroup.Get("/", eventHandler.FindAllEvents)
	eventGroup.Get("/:id", eventHandler.FindEventByID)
	eventGroup.Post("/", eventHandler.CreateEvent)
	eventGroup.Patch("/:id", eventHandler.PatchEvent)
	eventGroup.Delete("/:id", eventHandler.DeleteEvent)

	// Tag routes
	tagGroup := api.Group("/tags")
	tagGroup.Get("/", tagHandler.FindAllTags)
	tagGroup.Get("/:id", tagHandler.FindTagByID)
	tagGroup.Post("/", tagHandler.CreateTag)
	tagGroup.Patch("/:id", tagHandler.PatchTag)
	tagGroup.Delete("/:id", tagHandler.DeleteTag)

	// EventTag routes
	eventTagGroup := api.Group("/event_tags")
	eventTagGroup.Get("/", eventTagHandler.FindAllEventTags)
	eventTagGroup.Get("/event/:event_id", eventTagHandler.FindByEventID)
	eventTagGroup.Get("/tag/:tag_id", eventTagHandler.FindByTagID)
	eventTagGroup.Get("/:event_id/:tag_id", eventTagHandler.FindByEventAndTagID)
	eventTagGroup.Post("/", eventTagHandler.CreateEventTag)
	eventTagGroup.Delete("/:event_id/:tag_id", eventTagHandler.DeleteEventTag)
}
