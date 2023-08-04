package controller

import "github.com/hulhay/nagasari/models"

func buildGetMenusResponse(store *models.Store, menus []models.Menu) models.GetMenusResponse {
	result := models.GetMenusResponse{
		Store: buildStoreGetMenusResponse(store),
		Menu:  buildMenuGetMenusResponse(menus),
	}

	return result
}

func buildStoreGetMenusResponse(store *models.Store) models.StoreGetMenusReponse {
	result := models.StoreGetMenusReponse{
		UUID:             store.UUID,
		StoreName:        store.StoreName,
		OwnerUUID:        store.OwnerUUID,
		OwnerName:        store.OwnerName,
		OwnerPhoneNumber: store.OwnerPhoneNumber,
	}
	if store.StorePhotoURL != nil {
		result.StorePhotoURL = *store.StorePhotoURL
	}
	return result
}

func buildMenuGetMenusResponse(menus []models.Menu) []models.MenuGetMenusResponse {
	result := make([]models.MenuGetMenusResponse, len(menus))
	for i, menu := range menus {
		result[i] = buildMenu(menu)
	}
	return result
}

func buildMenu(menu models.Menu) models.MenuGetMenusResponse {
	return models.MenuGetMenusResponse{
		UUID:      menu.UUID,
		Name:      menu.Name,
		Price:     menu.Price,
		IsSoldOut: menu.IsSoldOut,
	}
}
