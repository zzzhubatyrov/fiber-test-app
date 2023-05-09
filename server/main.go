package main

import (
	"fmt"
	"log"
	"time"

	"fiber-test-app/internal/handlers/todoHandler"
	"fiber-test-app/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/create-models", func(c *fiber.Ctx) error {
		go func() {
			time.Sleep(1 * time.Second)
			todos := []models.Todo{}
			todoHandler.CreateModel(todos)
		}()
		if err := c.Redirect("/", fiber.StatusMovedPermanently); err != nil {
			return err
		}
		return nil
	})

	app.Get("/check-todo", func(c *fiber.Ctx) error {
		resultTodo := make(chan []models.Todo)
		var req models.Todo
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		go func() {
			time.Sleep(1 * time.Second)
			todos, _ := todoHandler.CheckTodo([]models.Todo{
				{
					Title:     req.Title,
					Completed: req.Completed,
				},
			})
			resultTodo <- todos
		}()
		return c.JSON(<-resultTodo)
	})

	app.Post("/create-todo", func(c *fiber.Ctx) error {
		resultChan := make(chan []models.Todo)
		var req models.Todo
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		go func() {
			time.Sleep(1 * time.Second)
			createTodo, _ := todoHandler.CreateTodo([]models.Todo{
				{
					Title:     req.Title,
					Completed: req.Completed,
				},
			})
			resultChan <- createTodo
		}()
		return c.JSON(<-resultChan)
	})

	fmt.Println("Starting server on port 5000...")
	log.Fatal(app.Listen(":5000"))
}
