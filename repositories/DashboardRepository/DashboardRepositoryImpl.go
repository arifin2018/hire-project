package dashboardrepository

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/azureOneSmile.git/db/models/DashboardModel"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	"gorm.io/gorm"
)

func (dashboardRepositoryImpl *DashboardRepositoryImpl) TicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB) (dashboards []dashboardmodel.DashboardModel, err error) {
	query := `
			SELECT 
				mp.id as AssigneeID,
				CONCAT_WS(' ', COALESCE(mp.firstName, ''), COALESCE(mp.middleName, ''), COALESCE(mp.lastName, '')) AS nameAssigned,
				SUM(CASE WHEN t.typeId = 2 AND t.isExternal = 1 THEN 1 ELSE 0 END) AS ExternalBugType2,
				SUM(CASE WHEN t.typeId = 1 AND t.isExternal = 1 THEN 1 ELSE 0 END) AS ExternalSupportType1,
				SUM(CASE WHEN t.typeId = 2 AND t.isExternal = 0 THEN 1 ELSE 0 END) AS InternalBugType2,
				SUM(CASE WHEN t.typeId = 1 AND t.isExternal = 0 THEN 1 ELSE 0 END) AS InternalSupportType1,
				SUM(CASE WHEN t.typeId = 5 THEN 1 ELSE 0 END) AS IOW,
				SUM(
					CASE 
						WHEN t.typeId = 2 AND t.isExternal = 1 THEN 1 
						WHEN t.typeId = 1 AND t.isExternal = 1 THEN 1
						WHEN t.typeId = 2 AND t.isExternal = 0 THEN 1
						WHEN t.typeId = 1 AND t.isExternal = 0 THEN 1
						WHEN t.typeId = 5 THEN 1
						ELSE 0 
					END
				) AS TotalTickets
			FROM Ticket t
			JOIN TicketMember tm ON t.id = tm.ticketId AND tm.isPIC = 1
			JOIN MemberPersonal mp ON tm.memberId = mp.id
			WHERE t.isDeleted IS NULL 
			AND t.typeId IN (1, 2, 5)
			GROUP BY mp.id, nameAssigned
			HAVING SUM(
				CASE 
					WHEN t.typeId = 2 AND t.isExternal = 1 THEN 1 
					WHEN t.typeId = 1 AND t.isExternal = 1 THEN 1
					WHEN t.typeId = 2 AND t.isExternal = 0 THEN 1
					WHEN t.typeId = 1 AND t.isExternal = 0 THEN 1
					WHEN t.typeId = 5 THEN 1
					ELSE 0 
				END
			) > 0
			ORDER BY nameAssigned;
		`
	if err := db.Raw(query).Scan(&dashboards).Error; err != nil {
		return dashboards, err
	}
	return dashboards, nil
}
func (dashboardRepositoryImpl *DashboardRepositoryImpl) ModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB) error {
	fmt.Println("user")
	// fmt.Printf("%+v\n", user)
	// if err := dashboardRepositoryImpl.DB.Create(user).Error; err != nil {
	// 	return err
	// }
	return nil
}
func (dashboardRepositoryImpl *DashboardRepositoryImpl) SubModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB) error {
	fmt.Println("user")
	// fmt.Printf("%+v\n", user)
	// if err := dashboardRepositoryImpl.DB.Create(user).Error; err != nil {
	// 	return err
	// }
	return nil
}
func (dashboardRepositoryImpl *DashboardRepositoryImpl) Create(app *fiber.Ctx, user *usermodel.User) error {
	fmt.Println("user")
	fmt.Printf("%+v\n", user)
	// if err := dashboardRepositoryImpl.DB.Create(user).Error; err != nil {
	// 	return err
	// }
	return nil
}
