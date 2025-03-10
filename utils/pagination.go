package utils

import (
	"strconv"
	"tasks-api/internal/entity"
)

const (
	DefaultPageQuery  = "1"
	DefaultLimitQuery = "10"
)

func PaginationFilterByQueryParams(page, limit string) (*entity.PaginationFilter, error) {
	if page == "" {
		page = DefaultPageQuery
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	if limit == "" {
		limit = DefaultLimitQuery
	}
	l, err := strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}

	return &entity.PaginationFilter{
		Page:  p,
		Limit: l,
	}, nil
}
