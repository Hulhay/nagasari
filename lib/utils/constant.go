package utils

import (
	"net/http"
	"regexp"

	"github.com/go-openapi/errors"
)

const (
	TimeFormat string = "2006-01-02 15:04:05"
	TimeZone   string = "Asia/Jakarta"
)

var (
	// 400
	ErrMalformatRequest = errors.New(http.StatusBadRequest, "malformat request")
	ErrKeywordLT3       = errors.New(http.StatusBadRequest, "at least 3 characters")
	ErrInvalidUUID      = errors.New(http.StatusBadRequest, "invalid uuid")
	ErrEmptyStoreUUID   = errors.New(http.StatusBadRequest, "store uuid cannot be empty")

	// 401
	ErrExpToken = errors.New(http.StatusUnauthorized, "token is expired")

	// 403
	ErrInvalidToken = errors.New(http.StatusForbidden, "invalid token")

	// 404
	ErrStoreNotFound = errors.New(http.StatusNotFound, "store not found")
	ErrEmailNotFound = errors.New(http.StatusNotFound, "email not found")

	// 422
	ErrInvalidEmail             = errors.New(http.StatusUnprocessableEntity, "invalid email format")
	ErrInvalidPhoneNumber       = errors.New(http.StatusUnprocessableEntity, "invalid phone number format")
	ErrEmptyEmail               = errors.New(http.StatusUnprocessableEntity, "email cannot be empty")
	ErrEmptyName                = errors.New(http.StatusUnprocessableEntity, "name cannot be empty")
	ErrEmptyPhoneNumber         = errors.New(http.StatusUnprocessableEntity, "phone number cannot be empty")
	ErrEmptyPassword            = errors.New(http.StatusUnprocessableEntity, "password cannot be empty")
	ErrShortPassword            = errors.New(http.StatusUnprocessableEntity, "password must have a minimum of 8 characters")
	ErrInvalidPassword          = errors.New(http.StatusUnprocessableEntity, "password must contain and contain only letters and numbers")
	ErrEmailAlreadyExists       = errors.New(http.StatusUnprocessableEntity, "email already exists")
	ErrPhoneNumberAlreadyExists = errors.New(http.StatusUnprocessableEntity, "phone number already exists")
	ErrWrongPassword            = errors.New(http.StatusUnprocessableEntity, "wrong password")

	// 500
	ErrWentWrong      = errors.New(http.StatusInternalServerError, "something went wrong")
	ErrFetchData      = errors.New(http.StatusInternalServerError, "an error occured when query db")
	ErrNoRowsAffected = errors.New(http.StatusInternalServerError, "no rows affected")
	ErrEncrypt        = errors.New(http.StatusInternalServerError, "something wrong with encryption process")
	ErrRegister       = errors.New(http.StatusInternalServerError, "failed to register")
)

var (
	PhoneNumberFormat = regexp.MustCompile(`(^0\d+$)`)
	EmailFormat       = regexp.MustCompile(`(^[\S]+@[\w]+\.[\w]{2,4}$|^[\S]+@[\w]+\.[\w]{2,4}\.[\w]{2}$)`)
)
