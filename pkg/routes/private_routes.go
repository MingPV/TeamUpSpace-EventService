package routes

import (
	userHandler "github.com/MingPV/EventService/internal/user/handler/rest"
	userRepository "github.com/MingPV/EventService/internal/user/repository"
	userUseCase "github.com/MingPV/EventService/internal/user/usecase"
	middleware "github.com/MingPV/EventService/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterPrivateRoutes(app fiber.Router, db *gorm.DB) {

	route := app.Group("/api/v1", middleware.JWTMiddleware())

	userRepo := userRepository.NewGormUserRepository(db)
	EventService := userUseCase.NewEventService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(EventService)

	route.Get("/me", userHandler.GetUser)

}
