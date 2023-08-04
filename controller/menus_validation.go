package controller

import (
	"github.com/google/uuid"
	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
)

func validateGetMenusRequest(req models.GetMenusRequest) error {
	if req.StoreUUID == "" {
		return utils.ErrEmptyStoreUUID
	}

	_, err := uuid.Parse(req.StoreUUID)
	if err != nil {
		return utils.ErrInvalidUUID
	}
	return nil
}
