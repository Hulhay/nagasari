package utils

import "math"

type Pagination struct {
	Page      int64   `json:"page,omitempty"`
	Size      int64   `json:"size,omitempty"`
	TotalData int64   `json:"total_data,omitempty"`
	TotalPage float64 `json:"total_page,omitempty"`
}

type PaginationRequest struct {
	Page       int64  `json:"page"`
	Limit      int64  `json:"limit"`
	Offset     int64  `json:"offset"`
	OrderField string `json:"orderBy"`
	Sort       string `json:"sort"`
}

func CalculatePagination(req PaginationRequest, total int) *Pagination {
	pagination := &Pagination{
		Page: req.Page,
		Size: req.Limit,
	}

	pagination.TotalData = int64(total)
	pagination.TotalPage = 0
	if req.Limit > 0 && req.Page > 0 {
		pagination.TotalPage = math.Ceil(float64(total) / float64(req.Limit))
	}

	return pagination
}
