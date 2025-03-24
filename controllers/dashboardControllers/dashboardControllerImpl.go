package dashboardcontrollers

import (
	"github.com/gofiber/fiber/v2"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	"github.com/lenna-ai/azureOneSmile.git/helpers"
)

func (dashboardControllerImpl *DashboardControllerImpl) TicketCompletionPerformace(app *fiber.Ctx) error {

	dashboard, err := dashboardControllerImpl.DashboardServices.TicketCompletionPerformace(app)
	if err != nil {
		return helpers.ResultFailedJsonApi(app, fiber.Map{}, err.Error())
	}

	result := fiber.Map{
		"data":   dashboard,
		"status": "successfully created",
	}

	return helpers.ResultSuccessCreateJsonApi(app, result)
}
func (dashboardControllerImpl *DashboardControllerImpl) ModalTicketCompletionPerformace(app *fiber.Ctx) error {
	DashboardModalTicketModel, err := dashboardControllerImpl.DashboardServices.ModalTicketCompletionPerformace(app)
	if err != nil {
		return helpers.ResultFailedJsonApi(app, fiber.Map{}, err.Error())
	}

	result := fiber.Map{
		"data":   DashboardModalTicketModel,
		"status": "successfully created",
	}

	return helpers.ResultSuccessCreateJsonApi(app, result)
}

func (dashboardControllerImpl *DashboardControllerImpl) SubModalTicketCompletionPerformace(app *fiber.Ctx) error {
	SubDashboardModalTicketModel, err := dashboardControllerImpl.DashboardServices.SubModalTicketCompletionPerformace(app)
	if err != nil {
		return helpers.ResultFailedJsonApi(app, fiber.Map{}, err.Error())
	}

	result := fiber.Map{
		"data":   SubDashboardModalTicketModel,
		"status": "successfully created",
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
