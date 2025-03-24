package dashboardrepository

import (
	"github.com/gofiber/fiber/v2"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
)

type DashboardRepository interface {
	Create(app *fiber.Ctx, user *usermodel.User) error
}

type DashboardRepositoryImpl struct {
}

func NewDashboardRepository() *DashboardRepositoryImpl {
	return &DashboardRepositoryImpl{}
}
