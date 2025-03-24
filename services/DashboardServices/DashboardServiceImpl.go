package dashboardservices

import (
	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/azureOneSmile.git/db/models/DashboardModel"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
)

func (dashboardServicesImpl *DashboardServicesImpl) TicketCompletionPerformace(app *fiber.Ctx) (dashboards []dashboardmodel.DashboardModel, err error) {
	db := dashboardServicesImpl.DB
	dashboards, err = dashboardServicesImpl.DashboardRepository.TicketCompletionPerformace(app, db)
	if err != nil {
		return dashboards, err
	}
	return dashboards, nil
}
func (dashboardServicesImpl *DashboardServicesImpl) ModalTicketCompletionPerformace(app *fiber.Ctx) error {
	db := dashboardServicesImpl.DB
	if err := dashboardServicesImpl.DashboardRepository.ModalTicketCompletionPerformace(app, db); err != nil {
		panic(err.Error())
	}
	return nil
}
func (dashboardServicesImpl *DashboardServicesImpl) SubModalTicketCompletionPerformace(app *fiber.Ctx) error {
	db := dashboardServicesImpl.DB
	if err := dashboardServicesImpl.DashboardRepository.SubModalTicketCompletionPerformace(app, db); err != nil {
		panic(err.Error())
	}
	return nil
}
func (dashboardServicesImpl *DashboardServicesImpl) Create(app *fiber.Ctx, user *usermodel.User) error {
	if err := dashboardServicesImpl.DashboardRepository.Create(app, user); err != nil {
		panic(err.Error())
	}
	return nil
}
