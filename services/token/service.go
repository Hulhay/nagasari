package token

import (
	"context"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/hulhay/nagasari/lib/config"
	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
)

type tokenService struct {
	cfg config.Config
}

type TokenService interface {
	GenerateToken(ctx context.Context, user *models.User) (*models.AuthResponse, error)
	ValidateToken(token string) (*jwt.Token, error)
}

func NewTokenService(cfg config.Config) TokenService {
	return &tokenService{
		cfg: cfg,
	}
}

func (ts *tokenService) GenerateToken(ctx context.Context, user *models.User) (*models.AuthResponse, error) {
	now := utils.Now()
	end := now.Add(ts.cfg.ExpJwtToken())
	claim := models.TokenClaim{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: end.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(ts.cfg.JWTSecretKey()))
	if err != nil {
		log.Printf("Error generate token: %v", err.Error())
		return nil, utils.ErrWentWrong
	}
	return &models.AuthResponse{
		Token:     signedToken,
		ExpiredAt: end.Format(utils.TimeFormat),
	}, nil
}

func (ts *tokenService) ValidateToken(token string) (*jwt.Token, error) {
	parse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.ErrInvalidToken
		}
		return []byte(ts.cfg.JWTSecretKey()), nil
	})
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "expired") {
			return nil, utils.ErrExpToken
		}
		log.Printf("Error validate token: %v", err.Error())
		return nil, utils.ErrWentWrong
	}
	return parse, nil
}
