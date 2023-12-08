package modules

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/m3rashid/htmx-wc/components"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		component := components.Index("Rashid")
		return component.Render(c.Context(), c.Response().BodyWriter())
	})

	app.Post("/clicked", func(c *fiber.Ctx) error {
		component := components.Hello("sdfkasdfaksdhfaksdf")
		return component.Render(context.Background(), c.Response().BodyWriter())
	})
}
