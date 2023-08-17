package auth

import (
	"context"
	"log"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
	"github.com/hulhay/nagasari/repositories"
)

type authService struct {
	repo repositories.Repository
}

type AuthService interface {
	Register(ctx context.Context, req models.RegisterRequest) error
}

func NewAuthService(repo repositories.Repository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (as *authService) Register(ctx context.Context, req models.RegisterRequest) error {
	err := as.validateRegisterRequest(ctx, req)
	if err != nil {
		return err
	}

	request, err := as.BuildRequest(req)
	if err != nil {
		return err
	}

	err = as.repo.InsertUserToDB(ctx, request)
	if err != nil {
		log.Printf("Error insert user to db : %v\n", err.Error())
		return utils.ErrRegister
	}

	return nil
}
