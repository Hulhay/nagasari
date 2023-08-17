package auth

import (
	"context"
	"log"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
	"github.com/hulhay/nagasari/repositories"
	ts "github.com/hulhay/nagasari/services/token"
)

type authService struct {
	repo         repositories.Repository
	tokenService ts.TokenService
}

type AuthService interface {
	Register(ctx context.Context, req models.RegisterRequest) error
	Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error)
}

func NewAuthService(
	repo repositories.Repository,
	tokenService ts.TokenService,
) AuthService {
	return &authService{
		repo:         repo,
		tokenService: tokenService,
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

func (as *authService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	user, err := as.validateLoginRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	err = utils.ValidatePassword(user.Password, req.Password)
	if err != nil {
		return nil, utils.ErrWrongPassword
	}

	response, err := as.tokenService.GenerateToken(ctx, user)
	if err != nil {
		return nil, err
	}

	return response, nil
}
