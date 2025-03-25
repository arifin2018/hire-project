package dashboardrepository

import (
	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/azureOneSmile.git/db/models/DashboardModel"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	"gorm.io/gorm"
)

type DashboardRepository interface {
	Create(app *fiber.Ctx, user *usermodel.User) error
	TicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB, pageSize int, page int) (dashboards []dashboardmodel.DashboardModel, err error)
	TotalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB) (totalCount int64, err error)
	ModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB, typeId int, isExternal int, assigneeId int) (DashboardModalTicketModel []dashboardmodel.DashboardModalTicketModel, err error)
	TotalModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB, typeId int, isExternal int) (totalCount int64, err error)
	SubModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB, typeId string, isExternal string, isPIC string) (SubDashboardModalTicketModel []dashboardmodel.DashboardSubModalTicketModel, err error)
}

type DashboardRepositoryImpl struct {
}

func NewDashboardRepository() *DashboardRepositoryImpl {
	return &DashboardRepositoryImpl{}
}
