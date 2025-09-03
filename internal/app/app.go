package app

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/MingPV/EventService/internal/entities"
	GrpcOrderHandler "github.com/MingPV/EventService/internal/order/handler/grpc"
	orderRepository "github.com/MingPV/EventService/internal/order/repository"
	orderUseCase "github.com/MingPV/EventService/internal/order/usecase"
	"github.com/MingPV/EventService/pkg/config"
	"github.com/MingPV/EventService/pkg/database"
	"github.com/MingPV/EventService/pkg/middleware"
	"github.com/MingPV/EventService/pkg/routes"
	orderpb "github.com/MingPV/EventService/proto/order"
)

// rest
func SetupRestServer(db *gorm.DB, cfg *config.Config) (*fiber.App, error) {
	app := fiber.New()
	middleware.FiberMiddleware(app)
	// comment out Swagger when testing
	routes.SwaggerRoute(app)
	routes.RegisterPublicRoutes(app, db)
	routes.RegisterPrivateRoutes(app, db)
	routes.RegisterNotFoundRoute(app)
	return app, nil
}

// grpc
func SetupGrpcServer(db *gorm.DB, cfg *config.Config) (*grpc.Server, error) {
	s := grpc.NewServer()
	orderRepo := orderRepository.NewGormOrderRepository(db)
	orderService := orderUseCase.NewOrderService(orderRepo)

	orderHandler := GrpcOrderHandler.NewGrpcOrderHandler(orderService)
	orderpb.RegisterOrderServiceServer(s, orderHandler)
	return s, nil
}

// dependencies
func SetupDependencies(env string) (*gorm.DB, *config.Config, error) {
	cfg := config.LoadConfig(env)

	db, err := database.Connect(cfg.DatabaseDSN)
	if err != nil {
		return nil, nil, err
	}

	if env == "test" {
		db.Migrator().DropTable(&entities.Order{})
	}
	// if err := db.AutoMigrate(&entities.Order{}); err != nil {
	if err := db.AutoMigrate(&entities.Event{}); err != nil {
		return nil, nil, err
	}

	return db, cfg, nil
}
