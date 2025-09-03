package main

import "github.com/MingPV/EventService/internal/app"

// @title TeamUpSpace Event Service API
// @version 1.0
// @description API documentation for the Event Service
// @host localhost:8003
// @BasePath /api/v1
func main() {
	app.Start() // Call server.go
}
