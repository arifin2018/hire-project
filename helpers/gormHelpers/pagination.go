package gormhelpers

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PaginatedResponseStruct struct {
	Data         interface{} `json:"data"`
	TotalCount   int64       `json:"total_all_data"`
	TotalPages   int         `json:"total_pages"`
	CurrentPage  int         `json:"current_page"`
	NextPage     int         `json:"next_page,omitempty"`
	PreviousPage int         `json:"previous_page,omitempty"`
}

func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		pageSize, _ := strconv.Atoi(c.Query("page_size"))

		if page == 0 && pageSize == 0 {
			return db
		}

		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func PaginatedResponse(page int, pageSize int, totalCount int64, data any) PaginatedResponseStruct {
	if page == 0 && pageSize == 0 {
		return PaginatedResponseStruct{
			Data:        data,
			TotalCount:  totalCount,
			TotalPages:  0,
			CurrentPage: page,
		}
	}

	if page <= 0 {
		page = 1
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	// Prepare the paginated response
	response := PaginatedResponseStruct{
		Data:        data,
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: page,
	}

	if page < totalPages {
		response.NextPage = page + 1
	}
	if page > 1 {
		response.PreviousPage = page - 1
	}

	return response
}
