package auth

import (
	"context"
	"log"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
)

func (as *authService) validateRegisterRequest(ctx context.Context, req models.RegisterRequest) error {
	user, err := as.repo.GetUserByEmailFromDB(ctx, req.Email)
	if err != nil {
		log.Printf("Error get user by email from db : %v\n", err.Error())
		return utils.ErrFetchData
	}
	if user != nil {
		return utils.ErrEmailAlreadyExists
	}

	user, err = as.repo.GetUserByPhoneNumberFromDB(ctx, req.PhoneNumber)
	if err != nil {
		log.Printf("Error get user by phone number from db : %v\n", err.Error())
		return utils.ErrFetchData
	}
	if user != nil {
		return utils.ErrPhoneNumberAlreadyExists
	}

	return nil
}
