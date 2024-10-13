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

	// Create
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

	// Update
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		// id, err := c.ParamsInt("id")  // ParamsInt for integer by default
		// id, err := c.ParamsInt("id")
		// if err != nil {
		// 	return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		// }

		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todo.Completed
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	log.Fatal(app.Listen(":3000"))
}
