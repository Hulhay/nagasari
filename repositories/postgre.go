package repositories

import (
	"context"

	"github.com/hulhay/nagasari/lib/config"
	"github.com/hulhay/nagasari/lib/database"
	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
)

type repository struct {
	qry database.SqlxDatabase
}

type Repository interface {
	GetStoresFromDB(ctx context.Context, req *models.GetStoresRequest) ([]models.Store, *utils.Pagination, error)
	GetStoreByStoreUUIDFromDB(ctx context.Context, storeUUID string) (*models.Store, error)

	GetProductsByStoreIDFromDB(ctx context.Context, storeID int64) ([]models.Product, error)

	GetUserByEmailFromDB(ctx context.Context, email string) (*models.User, error)
	GetUserByPhoneNumberFromDB(ctx context.Context, phoneNumber string) (*models.User, error)
	InsertUserToDB(ctx context.Context, req models.RegisterRequest) error
}

func NewRepository(cfg *config.Config) Repository {
	return &repository{
		qry: cfg.DB(),
	}
}
