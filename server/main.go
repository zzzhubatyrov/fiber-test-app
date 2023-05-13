package main

import (
	"fmt"
	"log"

	"fiber-test-app/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// GET requests
	app.Get("/", handlers.MainPage)
	app.Get("/check-todo", handlers.CheckTodo)

	// POST requests
	app.Post("/create-todo", handlers.CreateTodo)
	app.Post("/create-group", handlers.CreateGroup)

	// DELETE requests
	app.Delete("/delete-todo/:id", handlers.DeleteTodo)

	// PUT requests
	app.Put("/update-todo/:id", handlers.UpdateTodo)

	fmt.Println("Starting server on port 5000...")
	log.Fatal(app.Listen(":5000"))
}
