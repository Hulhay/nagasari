package repositories

import (
	"context"

	"github.com/hulhay/nagasari/models"
)

func (r *repository) GetProductsByStoreIDFromDB(ctx context.Context, storeID int64) ([]models.Product, error) {
	var menus []models.Product

	err := r.qry.Read().SelectContext(ctx, &menus, GetProductsByStoreIDQuery, storeID)
	if err != nil {
		return nil, err
	}

	return menus, nil
}
