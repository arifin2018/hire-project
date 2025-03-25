package dashboardservices

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/azureOneSmile.git/db/models/DashboardModel"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
)

func (dashboardServicesImpl *DashboardServicesImpl) TicketCompletionPerformace(app *fiber.Ctx, pageSize int, offset int) (dashboards []dashboardmodel.DashboardModel, totalCount int64, err error) {
	db := dashboardServicesImpl.DB
	var wg sync.WaitGroup
	var mu sync.Mutex              // Untuk menghindari race condition saat update variabel hasil
	errChan := make(chan error, 2) // Menyimpan error jika terjadi

	wg.Add(2)

	// Goroutine untuk mengambil data dashboards
	go func() {
		defer wg.Done()
		result, queryErr := dashboardServicesImpl.DashboardRepository.TicketCompletionPerformace(app, db, pageSize, offset)
		mu.Lock()
		if queryErr != nil {
			errChan <- queryErr
		} else {
			dashboards = result
		}
		mu.Unlock()
	}()

	// Goroutine untuk mengambil total count
	go func() {
		defer wg.Done()
		count, queryErr := dashboardServicesImpl.DashboardRepository.TotalTicketCompletionPerformace(app, db)
		mu.Lock()
		if queryErr != nil {
			errChan <- queryErr
		} else {
			totalCount = count
		}
		mu.Unlock()
	}()

	wg.Wait()
	close(errChan)

	// Cek apakah ada error dari salah satu query
	for queryErr := range errChan {
		if queryErr != nil {
			return dashboards, 0, queryErr
		}
	}

	return dashboards, totalCount, nil
}

func (dashboardServicesImpl *DashboardServicesImpl) ModalTicketCompletionPerformace(app *fiber.Ctx, pageSize int, page int, typeId int, isExternal int, assigneeId int) (DashboardModalTicketModel []dashboardmodel.DashboardModalTicketModel, totalCount int64, err error) {
	db := dashboardServicesImpl.DB
	DashboardModalTicketModel, err = dashboardServicesImpl.DashboardRepository.ModalTicketCompletionPerformace(app, db, pageSize, page, typeId, isExternal, assigneeId)
	if err != nil {
		return DashboardModalTicketModel, 0, err
	}
	count, err := dashboardServicesImpl.DashboardRepository.TotalModalTicketCompletionPerformace(app, db, typeId, isExternal, assigneeId)
	if err != nil {
		return DashboardModalTicketModel, 0, err
	}

	return DashboardModalTicketModel, count, nil
}

// SubModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB, typeId string, isExternal string, isPIC string) (SubDashboardModalTicketModel []dashboardmodel.DashboardSubModalTicketModel, err error)
func (dashboardServicesImpl *DashboardServicesImpl) SubModalTicketCompletionPerformace(app *fiber.Ctx) (SubDashboardModalTicketModel []dashboardmodel.DashboardSubModalTicketModel, err error) {
	db := dashboardServicesImpl.DB
	SubDashboardModalTicketModel, err = dashboardServicesImpl.DashboardRepository.SubModalTicketCompletionPerformace(app, db, 1, 1, 1, 1)
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
