package controller

import "github.com/hulhay/nagasari/models"

func buildGetMenusResponse(store *models.Store, products []models.Product) models.GetProductsResponse {
	result := models.GetProductsResponse{
		Store:   buildStoreGetProductsResponse(store),
		Product: buildProductGetProductsResponse(products),
	}

	return result
}

func buildStoreGetProductsResponse(store *models.Store) models.StoreGetProductsReponse {
	result := models.StoreGetProductsReponse{
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

func buildProductGetProductsResponse(products []models.Product) []models.ProductGetProductsResponse {
	result := make([]models.ProductGetProductsResponse, len(products))
	for i, product := range products {
		result[i] = buildProduct(product)
	}
	return result
}

func buildProduct(product models.Product) models.ProductGetProductsResponse {
	result := models.ProductGetProductsResponse{
		UUID:      product.UUID,
		Name:      product.Name,
		Price:     product.Price,
		IsSoldOut: product.IsSoldOut,
	}
	if product.ProductPhotoURL != nil {
		result.ProductPhotoURL = *product.ProductPhotoURL
	}
	return result
}
