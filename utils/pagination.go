package utils

import "tasks-api/internal/entity"

const (
	DefaultPage  = 1
	DefaultLimit = 10
)

func HandlePagination(filter *entity.PaginationFilter) {
	if filter.Page == 0 {
		filter.Page = DefaultPage
	}

	if filter.Limit == 0 {
		filter.Limit = DefaultLimit
	}
}
