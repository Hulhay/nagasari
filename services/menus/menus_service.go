package menus

import (
	"context"
	"log"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
	"github.com/hulhay/nagasari/repositories"
)

type menusService struct {
	repo repositories.Repository
}

type MenusService interface {
	GetMenusByStoreUUID(ctx context.Context, storeUUID string) (*models.Store, []models.Menu, error)
}

func NewStoresService(repo repositories.Repository) MenusService {
	return &menusService{
		repo: repo,
	}
}

func (ms *menusService) GetMenusByStoreUUID(ctx context.Context, storeUUID string) (*models.Store, []models.Menu, error) {
	store, err := ms.validateStoreUUID(ctx, storeUUID)
	if err != nil {
		return nil, nil, err
	}

	menus, err := ms.repo.GetMenusByStoreIDFromDB(ctx, store.ID)
	if err != nil {
		log.Printf("Error get stores from db : %v\n", err.Error())
		return nil, nil, utils.ErrFetchData
	}

	return store, menus, err
}
