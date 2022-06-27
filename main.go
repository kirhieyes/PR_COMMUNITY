package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug"
)

func main() {
	engine := pug.New("./views", ".pug")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("root call")
		return c.Render("index", fiber.Map{
			"Title": "Hello Index",
		})
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		fmt.Println("layout call")
		return c.Render("index", fiber.Map{
			"Title": "Hello layout",
		}, "layouts/main")
	})
	log.Fatal(app.Listen(":3000"))
}
