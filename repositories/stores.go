package repositories

import (
	"context"
	"database/sql"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
)

func (r *repository) GetStoresFromDB(ctx context.Context, req *models.GetStoresRequest) ([]models.Store, *utils.Pagination, error) {
	var (
		stores []models.Store
		err    error
		total  int
	)

	err = r.qry.Read().SelectContext(ctx, &stores, getStoresQuery, req.Keyword, req.Pagination.Limit, req.Pagination.Offset)
	if err != nil {
		return nil, nil, err
	}

	err = r.qry.Read().GetContext(ctx, &total, getStoresCountQuery, req.Keyword)
	if err != nil {
		return nil, nil, err
	}

	pagination := utils.CalculatePagination(req.Pagination, total)

	return stores, pagination, nil
}

func (r *repository) GetStoreByStoreUUIDFromDB(ctx context.Context, storeUUID string) (*models.Store, error) {
	var store models.Store

	err := r.qry.Read().GetContext(ctx, &store, getStoreByStoreUUIDQuery, storeUUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &store, nil
}
