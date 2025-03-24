package dashboardrepository

import (
	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/azureOneSmile.git/db/models/DashboardModel"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	"gorm.io/gorm"
)

type DashboardRepository interface {
	Create(app *fiber.Ctx, user *usermodel.User) error
	TicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB) (dashboards []dashboardmodel.DashboardModel, err error)
	ModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB) error
	SubModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB) error
}

type DashboardRepositoryImpl struct {
}

func NewDashboardRepository() *DashboardRepositoryImpl {
	return &DashboardRepositoryImpl{}
}
