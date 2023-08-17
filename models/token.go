package models

import "github.com/dgrijalva/jwt-go"

type TokenClaim struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}
