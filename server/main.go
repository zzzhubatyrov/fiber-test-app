package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
	ID   int
	Name string
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//est web-app on Fiber library [Go + Java]

	app.Get("/", func(c *fiber.Ctx) error {
		resultChan := make(chan string)

		go func() {
			time.Sleep(1 * time.Second)
			userData := map[int]string{
				1: "Hello World!",
			}
			usersJSON, err := json.Marshal(userData)
			if err != nil {
				resultChan <- err.Error()
				return
			}
			resultChan <- string(usersJSON)
		}()

		return c.JSON(<-resultChan)
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		resultChan := make(chan string)
		name := c.Params("name")

		userData := map[int]User{
			1: {ID: 1, Name: name},
		}

		// &User{
		// 	ID:   1,
		// 	Name: name,
		// }

		go func() {
			time.Sleep(1 * time.Second)
			usersJSON, err := json.Marshal(userData)
			if err != nil {
				resultChan <- err.Error()
				return
			}

			resultChan <- string(usersJSON)
		}()

		return c.JSON(<-resultChan)
	})

	fmt.Println("Starting server on port 5000...")
	log.Fatal(app.Listen(":5000"))
}
