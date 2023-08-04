package repositories

import (
	"context"

	"github.com/hulhay/nagasari/models"
)

func (r *repository) GetMenusByStoreIDFromDB(ctx context.Context, storeID int64) ([]models.Menu, error) {
	var menus []models.Menu

	err := r.qry.Read().SelectContext(ctx, &menus, GetMenusByStoreIDQuery, storeID)
	if err != nil {
		return nil, err
	}

	return menus, nil
}
