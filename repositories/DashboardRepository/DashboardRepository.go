package dashboardrepository

import (
	"github.com/gofiber/fiber/v2"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
)

type DashboardRepository interface {
	Create(app *fiber.Ctx, user *usermodel.User) error
	TicketCompletionPerformace(app *fiber.Ctx) error
	ModalTicketCompletionPerformace(app *fiber.Ctx) error
	SubModalTicketCompletionPerformace(app *fiber.Ctx) error
}

type DashboardRepositoryImpl struct {
}

func NewDashboardRepository() *DashboardRepositoryImpl {
	return &DashboardRepositoryImpl{}
}
