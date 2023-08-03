package stores

import (
	"context"
	"log"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
	"github.com/hulhay/nagasari/repositories"
)

type storesService struct {
	storesRepo repositories.Repository
}

type StoresService interface {
	GetStores(ctx context.Context, keyword string) ([]models.Store, error)
}

func NewStoresService(storesRepo repositories.Repository) StoresService {
	return &storesService{
		storesRepo: storesRepo,
	}
}

func (ss *storesService) GetStores(ctx context.Context, keyword string) ([]models.Store, error) {
	stores, err := ss.storesRepo.GetStoresFromDB(ctx, keyword)
	if err != nil {
		log.Printf("Error get stores from db : %v\n", err.Error())
		return nil, utils.ErrFetchData
	}
	return stores, nil
}
