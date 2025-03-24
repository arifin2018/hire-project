package dashboardmodel

type DashboardModel struct {
	AssigneeID           int    `gorm:"column:AssigneeID"`
	NameAssigned         string `gorm:"column:NameAssigned"`
	ExternalBugType2     int    `gorm:"column:External Bug (Type 2)"`
	ExternalSupportType1 int    `gorm:"column:External Support (Type 1)"`
	InternalBugType2     int    `gorm:"column:Internal Bug (Type 2)"`
	InternalSupportType1 int    `gorm:"column:Internal Support (Type 1)"`
	TotalTickets         int    `gorm:"column:TotalTickets"`
}
