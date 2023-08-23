package models

import "time"

type Product struct {
	ID              int64     `db:"id"`
	UUID            string    `db:"uuid"`
	Name            string    `db:"name"`
	Price           float64   `db:"price"`
	StoreID         string    `db:"store_id"`
	IsSoldOut       bool      `db:"is_sold_out"`
	ProductPhotoURL *string   `db:"product_photo_url"`
	CreatedAt       time.Time `db:"created_at"`
}

type GetProductsRequest struct {
	StoreUUID string
}

type GetProductsResponse struct {
	Store   StoreGetProductsReponse      `json:"store"`
	Product []ProductGetProductsResponse `json:"product"`
}

type StoreGetProductsReponse struct {
	UUID             string `json:"uuid"`
	StoreName        string `json:"store_name"`
	StorePhotoURL    string `json:"store_photo_url"`
	OwnerUUID        string `json:"owner_uuid"`
	OwnerName        string `json:"owner_name"`
	OwnerPhoneNumber string `json:"owner_phone_number"`
}

type ProductGetProductsResponse struct {
	UUID            string  `json:"uuid"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	IsSoldOut       bool    `json:"is_sold_out"`
	ProductPhotoURL string  `json:"product_photo_url"`
}
