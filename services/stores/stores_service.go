package stores

import (
	"context"
	"log"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
	"github.com/hulhay/nagasari/repositories"
)

type storesService struct {
	repo repositories.Repository
}

type StoresService interface {
	GetStores(ctx context.Context, req *models.GetStoresRequest) ([]models.Store, *utils.Pagination, error)
}

func NewStoresService(repo repositories.Repository) StoresService {
	return &storesService{
		repo: repo,
	}
}

func (ss *storesService) GetStores(ctx context.Context, req *models.GetStoresRequest) ([]models.Store, *utils.Pagination, error) {
	stores, pagination, err := ss.repo.GetStoresFromDB(ctx, req)
	if err != nil {
		log.Printf("Error get stores from db : %v\n", err.Error())
		return nil, nil, utils.ErrFetchData
	}
	return stores, pagination, err
}
