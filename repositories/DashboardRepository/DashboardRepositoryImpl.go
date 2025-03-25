package dashboardrepository

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/azureOneSmile.git/db/models/DashboardModel"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	gormhelpers "github.com/lenna-ai/azureOneSmile.git/helpers/gormHelpers"
	"gorm.io/gorm"
)

func (dashboardRepositoryImpl *DashboardRepositoryImpl) TicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB, pageSize int, page int) (dashboards []dashboardmodel.DashboardModel, err error) {
	query := `
			SELECT 
				mp.id as AssigneeID,
				CONCAT_WS(' ', COALESCE(mp.firstName, ''), COALESCE(mp.middleName, ''), COALESCE(mp.lastName, '')) AS NameAssigned,
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
			ORDER BY nameAssigned
			LIMIT ? OFFSET ?;
		`
	if err := db.Raw(query, pageSize, page).Scan(&dashboards).Error; err != nil {
		return dashboards, err
	}
	return dashboards, nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) TotalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB) (totalCount int64, err error) {
	query := fmt.Sprintf(`
			SELECT COUNT(*) FROM (
				SELECT 
					mp.id as AssigneeID
				FROM Ticket t
				JOIN TicketMember tm ON t.id = tm.ticketId AND tm.isPIC = 1
				JOIN MemberPersonal mp ON tm.memberId = mp.id
				WHERE t.isDeleted IS NULL 
				AND t.typeId IN (1, 2, 5)
				GROUP BY mp.id
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
			) AS paginatedResult;
		`)
	if err := db.Raw(query).Scan(&totalCount).Error; err != nil {
		return totalCount, err
	}
	return totalCount, nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) ModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB, pageSize int, page int, typeId int, isExternal int, assigneeId int) (DashboardModalTicketModel []dashboardmodel.DashboardModalTicketModel, err error) {
	query := fmt.Sprintf(`
			SELECT 
				t.documentNo AS TicketDocumentNo,
				t.estimatedManhours AS EstimatedManhours
			FROM Ticket t
			JOIN (
				SELECT ticketId, MIN(memberId) AS memberId 
				FROM TicketMember 
				WHERE isPIC = 1 
				GROUP BY ticketId
			) tm ON t.id = tm.ticketId
			JOIN MemberPersonal mp ON tm.memberId = mp.id
			WHERE t.isDeleted IS NULL 
			AND t.typeId = %v 
			AND t.isExternal = %v
			AND mp.id = %v
			ORDER BY t.createdAt DESC
			LIMIT ? OFFSET ?;
	`, typeId, isExternal, assigneeId)
	if err := db.Raw(query, pageSize, page).Scan(&DashboardModalTicketModel).Error; err != nil {
		return DashboardModalTicketModel, err
	}
	return DashboardModalTicketModel, nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) TotalModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB, typeId int, isExternal int, assigneeId int) (totalCount int64, err error) {
	query := fmt.Sprintf(`
		SELECT COUNT(*) AS TotalCount
			FROM (
				SELECT 
					t.documentNo AS TicketDocumentNo,
					t.estimatedManhours AS EstimatedManhours
				FROM Ticket t
				JOIN (
					SELECT ticketId, MIN(memberId) AS memberId 
					FROM TicketMember 
					WHERE isPIC = 1 
					GROUP BY ticketId
				) tm ON t.id = tm.ticketId
				JOIN MemberPersonal mp ON tm.memberId = mp.id
				WHERE t.isDeleted IS NULL 
				AND t.typeId = %v 
				AND t.isExternal = %v
				AND mp.id = %v
			) AS subquery;
	`, typeId, isExternal, assigneeId)
	if err := db.Raw(query).Scan(&totalCount).Error; err != nil {
		return totalCount, err
	}
	return totalCount, nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) SubModalTicketCompletionPerformace(app *fiber.Ctx, db *gorm.DB, typeId int, isExternal int, isPIC int, assigneeId int) (SubDashboardModalTicketModel []dashboardmodel.DashboardSubModalTicketModel, err error) {
	query := fmt.Sprintf(`
		SELECT 
			IFNULL(SUM(t.estimatedManhours), 0) AS totalEstimatedWork
		FROM Ticket t
		JOIN TicketMember tm ON t.id = tm.ticketId 
		JOIN MemberPersonal mp ON tm.memberId = mp.id
		WHERE t.isDeleted IS NULL 
		AND t.typeId = %v
		AND t.isExternal = %v
		AND tm.isPIC = %v  
		AND mp.id = %v;
	`, typeId, isExternal, isPIC, assigneeId)
	if err := db.Scopes(gormhelpers.Paginate(app)).Raw(query).Scan(&SubDashboardModalTicketModel).Error; err != nil {
		return SubDashboardModalTicketModel, err
	}
	return SubDashboardModalTicketModel, nil
}
func (dashboardRepositoryImpl *DashboardRepositoryImpl) Create(app *fiber.Ctx, user *usermodel.User) error {
	fmt.Println("user")
	fmt.Printf("%+v\n", user)
	// if err := dashboardRepositoryImpl.DB.Create(user).Error; err != nil {
	// 	return err
	// }
	return nil
}
