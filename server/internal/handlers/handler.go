package handlers

import (
	"fiber-test-app/internal/handlers/groupHandlers"
	"fiber-test-app/internal/handlers/todoHandlers"
	"fiber-test-app/internal/models"
	"log"
	"strconv"

	"github.com/gammazero/workerpool"
	"github.com/gofiber/fiber/v2"
)

// handler link
var pool = workerpool.New(10)

func MainPage(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func CheckTodo(c *fiber.Ctx) error {
	var req models.Todo

	resultTodo := make(chan []models.Todo)
	pool.Submit(func() {
		checkTodo, _ := todoHandlers.CheckTodo([]models.Todo{
			{
				Title:     req.Title,
				Completed: req.Completed,
			},
		})
		resultTodo <- checkTodo
	})
	return c.JSON(<-resultTodo)
}

func CreateTodo(c *fiber.Ctx) error {
	resultChan := make(chan string)

	return c.JSON(<-resultChan)
}

func DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	resultChan := make(chan error)
	pool.Submit(func() {
		err := todoHandlers.DeleteTodo(uint(id))
		if err != nil {
			log.Println(err)
		}
		resultChan <- err
	})

	if err := <-resultChan; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var req models.Todo

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := todoHandlers.UpdateTodo(id, &req); err != nil {
		return err
	}

	return c.JSON(req)
}

func CreateGroup(c *fiber.Ctx) error {
	var req models.Group
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	resultChan := make(chan models.Group)
	pool.Submit(func() {
		createdGroup, err := groupHandlers.CreateGroup(req)
		if err != nil {
			resultChan <- models.Group{} // отправляем пустую структуру в канал в случае ошибки
			return
		}

		resultChan <- createdGroup
	})

	return c.JSON(<-resultChan)
}
