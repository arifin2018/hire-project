package dashboardcontrollers

import (
	"github.com/gofiber/fiber/v2"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	"github.com/lenna-ai/azureOneSmile.git/helpers"
)

func (dashboardControllerImpl *DashboardControllerImpl) TicketCompletionPerformace(app *fiber.Ctx) error {

	// if err := dashboardControllerImpl.DashboardServices.Create(app, user); err != nil {
	// 	return err
	// }

	result := fiber.Map{
		"data": "successfully created",
	}

	return helpers.ResultSuccessCreateJsonApi(app, result)
}
func (dashboardControllerImpl *DashboardControllerImpl) ModalTicketCompletionPerformace(app *fiber.Ctx) error {

	// if err := dashboardControllerImpl.DashboardServices.Create(app, user); err != nil {
	// 	return err
	// }

	result := fiber.Map{
		"data": "successfully created",
	}

	return helpers.ResultSuccessCreateJsonApi(app, result)
}
func (dashboardControllerImpl *DashboardControllerImpl) SubModalTicketCompletionPerformace(app *fiber.Ctx) error {

	// if err := dashboardControllerImpl.DashboardServices.Create(app, user); err != nil {
	// 	return err
	// }

	result := fiber.Map{
		"data": "successfully created",
	}

	return helpers.ResultSuccessCreateJsonApi(app, result)
}

func (dashboardControllerImpl *DashboardControllerImpl) Create(app *fiber.Ctx) error {
	user := new(usermodel.User)
	if err := app.BodyParser(user); err != nil {
		return err
	}
	if err := dashboardControllerImpl.DashboardServices.Create(app, user); err != nil {
		return err
	}

	result := fiber.Map{
		"data": "successfully created",
	}

	return helpers.ResultSuccessCreateJsonApi(app, result)
}
