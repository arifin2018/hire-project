package dashboardservices

import (
	"github.com/gofiber/fiber/v2"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
)

func (dashboardServicesImpl *DashboardServicesImpl) TicketCompletionPerformace(app *fiber.Ctx) error {
	if err := dashboardServicesImpl.DashboardRepository.TicketCompletionPerformace(app); err != nil {
		panic(err.Error())
	}
	return nil
}
func (dashboardServicesImpl *DashboardServicesImpl) ModalTicketCompletionPerformace(app *fiber.Ctx) error {
	if err := dashboardServicesImpl.DashboardRepository.ModalTicketCompletionPerformace(app); err != nil {
		panic(err.Error())
	}
	return nil
}
func (dashboardServicesImpl *DashboardServicesImpl) SubModalTicketCompletionPerformace(app *fiber.Ctx) error {
	if err := dashboardServicesImpl.DashboardRepository.SubModalTicketCompletionPerformace(app); err != nil {
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
