package dashboardcontrollers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	"github.com/lenna-ai/azureOneSmile.git/helpers"
	gormhelpers "github.com/lenna-ai/azureOneSmile.git/helpers/gormHelpers"
)

func (dashboardControllerImpl *DashboardControllerImpl) TicketCompletionPerformace(app *fiber.Ctx) error {
	page, _ := strconv.Atoi(app.Query("page"))
	pageSize, _ := strconv.Atoi(app.Query("page_size"))
	dashboard, totalCount, err := dashboardControllerImpl.DashboardServices.TicketCompletionPerformace(app, pageSize, page)
	if err != nil {
		return helpers.ResultFailedJsonApi(app, fiber.Map{}, err.Error())
	}

	return helpers.ResultSuccessJsonApi(app, gormhelpers.PaginatedResponse(page, pageSize, totalCount, dashboard))
}

func (dashboardControllerImpl *DashboardControllerImpl) ModalTicketCompletionPerformace(app *fiber.Ctx) error {
	page, _ := strconv.Atoi(app.Query("page"))
	pageSize, _ := strconv.Atoi(app.Query("page_size"))
	typeId, _ := strconv.Atoi(app.Query("typeId"))
	isExternal, _ := strconv.Atoi(app.Query("isExternal"))
	assigneeId, _ := strconv.Atoi(app.Query("assigneeId"))

	DashboardModalTicketModel, totalCount, err := dashboardControllerImpl.DashboardServices.ModalTicketCompletionPerformace(app, pageSize, page, typeId, isExternal, assigneeId)
	if err != nil {
		return helpers.ResultFailedJsonApi(app, fiber.Map{}, err.Error())
	}

	return helpers.ResultSuccessJsonApi(app, gormhelpers.PaginatedResponse(page, pageSize, totalCount, DashboardModalTicketModel))
}

func (dashboardControllerImpl *DashboardControllerImpl) SubModalTicketCompletionPerformace(app *fiber.Ctx) error {
	page, _ := strconv.Atoi(app.Query("page"))
	pageSize, _ := strconv.Atoi(app.Query("page_size"))
	typeId, _ := strconv.Atoi(app.Query("typeId"))
	isExternal, _ := strconv.Atoi(app.Query("isExternal"))
	isPIC, _ := strconv.Atoi(app.Query("isPIC"))
	assigneeId, _ := strconv.Atoi(app.Query("assigneeId"))
	SubDashboardModalTicketModel, totalCount, err := dashboardControllerImpl.DashboardServices.SubModalTicketCompletionPerformace(app, pageSize, page, typeId, isExternal, isPIC, assigneeId)
	if err != nil {
		return helpers.ResultFailedJsonApi(app, fiber.Map{}, err.Error())
	}
	return helpers.ResultSuccessJsonApi(app, gormhelpers.PaginatedResponse(page, pageSize, totalCount, SubDashboardModalTicketModel))
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
