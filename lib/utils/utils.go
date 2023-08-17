package utils

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}

func ValidatePassword(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return err
	}
	return nil
}

func ValidateEmailFormat(email string) error {
	if ok := EmailFormat.MatchString(email); !ok {
		return ErrInvalidEmail
	}
	return nil
}

func ValidatePhoneNumberFormat(phoneNumber string) error {
	if ok := PhoneNumberFormat.MatchString(phoneNumber); !ok {
		return ErrInvalidPhoneNumber
	}
	return nil
}

func ValidatePasswordFormat(password string) error {
	var letter, number, special bool

	if len(password) < 8 {
		return ErrShortPassword
	}

	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			number = true
		case unicode.IsLetter(char):
			letter = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		}
	}
	if !number || !letter || special {
		return ErrInvalidPassword
	}

	return nil
}
