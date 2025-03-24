package dashboardservices

import (
	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/azureOneSmile.git/db/models/DashboardModel"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	dashboardrepository "github.com/lenna-ai/azureOneSmile.git/repositories/DashboardRepository"
	"gorm.io/gorm"
)

type DashboardServices interface {
	Create(app *fiber.Ctx, user *usermodel.User) error
	TicketCompletionPerformace(app *fiber.Ctx) (dashboards []dashboardmodel.DashboardModel, err error)
	ModalTicketCompletionPerformace(app *fiber.Ctx) error
	SubModalTicketCompletionPerformace(app *fiber.Ctx) error
}

type DashboardServicesImpl struct {
	DashboardRepository dashboardrepository.DashboardRepository
	DB                  *gorm.DB
}

func NewDashboardServices(dashboardRepository dashboardrepository.DashboardRepository, db *gorm.DB) *DashboardServicesImpl {
	return &DashboardServicesImpl{
		DashboardRepository: dashboardRepository,
		DB:                  db,
	}
}
