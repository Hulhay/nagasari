package utils

import (
	"net/http"

	"github.com/go-openapi/errors"
)

const (
	TimeFormat string = "2006-01-02 15:04:05"
	TimeZone   string = "Asia/Jakarta"
)

var (
	ErrMalformatRequest = errors.New(http.StatusBadRequest, "malformat request")
	ErrKeywordLT3       = errors.New(http.StatusBadRequest, "at least 3 characters")

	ErrFetchData = errors.New(http.StatusInternalServerError, "an error occured when query db")
)
