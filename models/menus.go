package models

import "time"

type Menu struct {
	ID        int64     `db:"id"`
	UUID      string    `db:"uuid"`
	Name      string    `db:"name"`
	Price     float64   `db:"price"`
	StoreID   string    `db:"store_id"`
	IsSoldOut bool      `db:"is_sold_out"`
	CreatedAt time.Time `db:"created_at"`
}

type GetMenusRequest struct {
	StoreUUID string
}

type GetMenusResponse struct {
	Store StoreGetMenusReponse   `json:"store"`
	Menu  []MenuGetMenusResponse `json:"menu"`
}

type StoreGetMenusReponse struct {
	UUID             string `json:"uuid"`
	StoreName        string `json:"store_name"`
	StorePhotoURL    string `json:"store_photo_url"`
	OwnerUUID        string `json:"owner_uuid"`
	OwnerName        string `json:"owner_name"`
	OwnerPhoneNumber string `json:"owner_phone_number"`
}

type MenuGetMenusResponse struct {
	UUID      string  `json:"uuid"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	IsSoldOut bool    `json:"is_sold_out"`
}
