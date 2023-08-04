package controller

import (
	"strconv"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
)

func validateGetStoresRequest(req models.GetStoresRequest) (*models.GetStoresRequest, error) {
	var err error

	if req.Page != "" {
		req.Pagination.Page, err = strconv.ParseInt(req.Page, 10, 64)
		if err != nil {
			return nil, utils.ErrMalformatRequest
		}
	}

	if req.Pagination.Page < 1 {
		req.Pagination.Page = 1
	}

	if req.Size != "" {
		req.Pagination.Limit, err = strconv.ParseInt(req.Size, 10, 64)
		if err != nil {
			return nil, utils.ErrMalformatRequest
		}
	}

	if req.Pagination.Limit < 1 || req.Pagination.Limit > 20 {
		req.Pagination.Limit = 20
	}

	req.Pagination.Offset = (req.Pagination.Limit * req.Pagination.Page) - req.Pagination.Limit

	if req.Keyword != "" && len(req.Keyword) < 3 {
		return nil, utils.ErrKeywordLT3
	}

	return &req, nil
}
