package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Text      string `json:"text"`
}

func main() {
	fmt.Println("Hello, World!")

	// var name = "John"
	// const name2 = "Doe"

	// fmt.Println(name)
	// fmt.Println(name2)

	// name3 := "Blue Eyes White Dragon"
	// fmt.Println(name3)

	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, World!")
		return c.Status(200).JSON(fiber.Map{"msg": "Blue Eyes White Dragon!"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		if todo.Text == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Text is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todos)
	})

	log.Fatal(app.Listen(":3000"))
}
