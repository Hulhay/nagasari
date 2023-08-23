package products

import (
	"context"
	"log"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
	"github.com/hulhay/nagasari/repositories"
)

type productsService struct {
	repo repositories.Repository
}

type ProductsService interface {
	GetProductsByStoreUUID(ctx context.Context, storeUUID string) (*models.Store, []models.Product, error)
}

func NewProductsService(repo repositories.Repository) ProductsService {
	return &productsService{
		repo: repo,
	}
}

func (ps *productsService) GetProductsByStoreUUID(ctx context.Context, storeUUID string) (*models.Store, []models.Product, error) {
	store, err := ps.validateStoreUUID(ctx, storeUUID)
	if err != nil {
		return nil, nil, err
	}

	menus, err := ps.repo.GetProductsByStoreIDFromDB(ctx, store.ID)
	if err != nil {
		log.Printf("Error get stores from db : %v\n", err.Error())
		return nil, nil, utils.ErrFetchData
	}

	return store, menus, err
}
