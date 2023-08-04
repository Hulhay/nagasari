package menus

import (
	"context"
	"log"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
)

func (ms *menusService) validateStoreUUID(ctx context.Context, storeUUID string) (*models.Store, error) {
	store, err := ms.repo.GetStoreByStoreUUIDFromDB(ctx, storeUUID)
	if err != nil {
		log.Printf("Error get stores from db : %v\n", err.Error())
		return nil, utils.ErrFetchData
	}
	if store == nil {
		return nil, utils.ErrStoreNotFound
	}

	return store, nil
}
