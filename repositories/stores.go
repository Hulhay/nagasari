package repositories

import (
	"context"
	"fmt"

	"github.com/hulhay/nagasari/models"
)

func (r *repository) GetStoresFromDB(ctx context.Context, keyword string) ([]models.Store, error) {
	var (
		results []models.Store
		args    []interface{}
	)

	basicQuery := getStoresQuery
	if keyword != "" {
		basicQuery = fmt.Sprintf("%s WHERE %s", basicQuery, getStoresByKeywordName)
		args = append(args, keyword)
	}

	stmt, err := r.qry.Read().PrepareContext(ctx, basicQuery)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var result models.Store
		err := rows.Scan(&result.ID, &result.UUID, &result.StoreName, &result.StorePhotoURL, &result.OwnerUUID, &result.OwnerName, &result.OwnerPhoneNumber, &result.CreatedAt)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}
