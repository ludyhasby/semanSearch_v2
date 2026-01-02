package route

import (
	"sahih-be/internal/delivery"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Selamat Datang di Backend Sahih")
	})
	// Route Doa
	doaRoutes := app.Group("Doa")
	doaRoutes.Get("/", delivery.GetAllDoa)
}
