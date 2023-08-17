package auth

import (
	"log"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
)

func (as *authService) BuildRequest(req models.RegisterRequest) (models.RegisterRequest, error) {
	encryptedPassword, err := utils.EncryptPassword(req.Password)
	if err != nil {
		log.Printf("Error encrypt password : %v\n", err.Error())
		return models.RegisterRequest{}, utils.ErrEncrypt
	}

	req.EncryptedPassword = encryptedPassword

	return req, err
}
