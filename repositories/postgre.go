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
}

func NewRepository(cfg *config.Config) Repository {
	return &repository{
		qry: cfg.DB(),
	}
}
