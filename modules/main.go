package modules

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": os.Getenv("APP_NAME"),
			// "headerItems": fiber.Map{
			// 	"Dashboard": "/dashboard",
			// 	"Team":      "/team",
			// 	"Projects":  "/projects",
			// 	"Calender":  "/calender",
			// },
			// "activeHeader": "/",
		})
	})

	app.Post("/clicked", func(c *fiber.Ctx) error {
		return c.SendString("<fluent-button id='parent-div'>Clicked hello world</fluent-button>")
	})
}
