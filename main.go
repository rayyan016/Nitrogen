package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello, World!")

	// var name = "John"
	// const name2 = "Doe"

	// fmt.Println(name)
	// fmt.Println(name2)

	// name3 := "Blue Eyes White Dragon"
	// fmt.Println(name3)

	app := fiber.New()

	log.Fatal(app.Listen(":3000"))
}
