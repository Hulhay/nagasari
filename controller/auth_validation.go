package controller

import (
	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
)

func (ac *authController) validateRegisterRequest(req models.RegisterRequest) error {
	var err error

	if req.Name == `` {
		return utils.ErrEmptyName
	}

	if req.Email == `` {
		return utils.ErrEmptyEmail
	}
	err = utils.ValidateEmailFormat(req.Email)
	if err != nil {
		return err
	}

	if req.PhoneNumber == `` {
		return utils.ErrEmptyPhoneNumber
	}
	err = utils.ValidatePhoneNumberFormat(req.PhoneNumber)
	if err != nil {
		return err
	}

	if req.Password == `` {
		return utils.ErrEmptyPassword
	}
	err = utils.ValidatePasswordFormat(req.Password)
	if err != nil {
		return err
	}

	return nil
}

func (ac *authController) validateLoginRequest(req models.LoginRequest) error {
	var err error

	if req.Email == `` {
		return utils.ErrEmptyEmail
	}
	err = utils.ValidateEmailFormat(req.Email)
	if err != nil {
		return err
	}

	if req.Password == `` {
		return utils.ErrEmptyPassword
	}

	return nil
}
