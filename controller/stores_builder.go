package controller

import (
	"github.com/hulhay/nagasari/models"
)

func buildGetStoresResponse(stores []models.Store) []models.GetStoresResponse {
	result := make([]models.GetStoresResponse, len(stores))
	for i, store := range stores {
		result[i] = buildGetStores(store)
	}
	return result
}

func buildGetStores(store models.Store) models.GetStoresResponse {
	result := models.GetStoresResponse{
		UUID:      store.UUID,
		StoreName: store.StoreName,
	}
	if store.StorePhotoURL != nil {
		result.StorePhotoURL = *store.StorePhotoURL
	}

	return result
}
