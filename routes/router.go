package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/azureOneSmile.git/injector"
)

func Router(app *fiber.App) {
	allController := injector.InitializeController()
	// localhost:3000/dashboard/ticketCompletionPerformace
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"data":   nil,
			"status": "ini hanya test dari arifin",
		})
	})
	dashboard := app.Group("dashboard")
	// ticket completion performance
	dashboard.Post("/", allController.DashboardController.Create)

	ticketCompletionPerformace := app.Group("ticketCompletionPerformace")
	ticketCompletionPerformace.Get("/", allController.DashboardController.TicketCompletionPerformace)
	ticketCompletionPerformace.Get("/modal", allController.DashboardController.ModalTicketCompletionPerformace)
	ticketCompletionPerformace.Get("/submodal", allController.DashboardController.SubModalTicketCompletionPerformace)
}
