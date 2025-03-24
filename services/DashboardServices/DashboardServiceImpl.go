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
func (dashboardServicesImpl *DashboardServicesImpl) ModalTicketCompletionPerformace(app *fiber.Ctx) (DashboardModalTicketModel []dashboardmodel.DashboardModalTicketModel, err error) {
	db := dashboardServicesImpl.DB
	DashboardModalTicketModel, err = dashboardServicesImpl.DashboardRepository.ModalTicketCompletionPerformace(app, db, "1", "1")
	if err != nil {
		return DashboardModalTicketModel, err
	}

	return DashboardModalTicketModel, nil
}

// SubModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB, typeId string, isExternal string, isPIC string) (SubDashboardModalTicketModel []dashboardmodel.DashboardSubModalTicketModel, err error)
func (dashboardServicesImpl *DashboardServicesImpl) SubModalTicketCompletionPerformace(app *fiber.Ctx) (SubDashboardModalTicketModel []dashboardmodel.DashboardSubModalTicketModel, err error) {
	db := dashboardServicesImpl.DB
	SubDashboardModalTicketModel, err = dashboardServicesImpl.DashboardRepository.SubModalTicketCompletionPerformace(app, db, "1", "1", "1")
	if err != nil {
		return SubDashboardModalTicketModel, err
	}
	return SubDashboardModalTicketModel, err
}

func (dashboardServicesImpl *DashboardServicesImpl) Create(app *fiber.Ctx, user *usermodel.User) error {
	if err := dashboardServicesImpl.DashboardRepository.Create(app, user); err != nil {
		panic(err.Error())
	}
	return nil
}
