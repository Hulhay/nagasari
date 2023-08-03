package utils

import (
	"net/http"

	"github.com/go-openapi/errors"
)

const (
	TimeFormat string = "2006-01-02 15:04:05.000-0700"
)

var (
	ErrKeywordLT3 = errors.New(http.StatusBadRequest, "at least 3 characters")

	ErrFetchData = errors.New(http.StatusInternalServerError, "an error occured when query db")
)
