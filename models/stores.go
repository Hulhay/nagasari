package models

import (
	"time"

	"github.com/hulhay/nagasari/lib/utils"
)

type Store struct {
	ID               int64     `db:"id"`
	UUID             string    `db:"uuid"`
	StoreName        string    `db:"store_name"`
	StorePhotoURL    *string   `db:"store_photo_url"`
	OwnerUUID        string    `db:"owner_uuid"`
	OwnerName        string    `db:"owner_name"`
	OwnerPhoneNumber string    `db:"owner_phone_number"`
	CreatedAt        time.Time `db:"created_at"`
}

type GetStoresRequest struct {
	Pagination utils.PaginationRequest
	Page       string
	Size       string
	Keyword    string
}

type GetStoresResponse struct {
	UUID             string `json:"uuid"`
	StoreName        string `json:"store_name"`
	StorePhotoURL    string `json:"store_photo_url"`
	OwnerUUID        string `json:"owner_uuid"`
	OwnerName        string `json:"owner_name"`
	OwnerPhoneNumber string `json:"owner_phone_number"`
}
