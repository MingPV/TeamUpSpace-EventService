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
)

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB) {

	api := app.Group("/api/v1")

	// === Dependency Wiring ===

	// // Order
	// orderRepo := orderRepository.NewGormOrderRepository(db)
	// orderService := orderUseCase.NewOrderService(orderRepo)
	// orderHandler := orderHandler.NewHttpOrderHandler(orderService)

	// Event
	eventRepo := eventRepository.NewGormEventRepository(db)
	eventService := eventUseCase.NewEventService(eventRepo)
	eventHandler := eventHandler.NewHttpEventHandler(eventService)

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
}
