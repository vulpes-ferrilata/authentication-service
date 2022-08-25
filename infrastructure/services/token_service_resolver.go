package services

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/config"
)

type tokenType string

var (
	AccessToken  tokenType = "access token"
	RefreshToken tokenType = "refresh token"
)

type TokenServiceResolver interface {
	GetTokenService(tokenType tokenType) TokenService
}

func NewTokenServiceResolver(config config.Config) (TokenServiceResolver, error) {
	accessTokenService, err := NewTokenService(config.AccessToken.Algorithm,
		config.AccessToken.SecretKey,
		config.AccessToken.Expiration)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	refreshTokenService, err := NewTokenService(config.RefreshToken.Algorithm,
		config.RefreshToken.SecretKey,
		config.RefreshToken.Expiration)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &tokenServiceResolver{
		accessTokenService:  accessTokenService,
		refreshTokenService: refreshTokenService,
	}, nil
}

type tokenServiceResolver struct {
	accessTokenService  TokenService
	refreshTokenService TokenService
}

func (t tokenServiceResolver) GetTokenService(tokenType tokenType) TokenService {
	switch tokenType {
	case AccessToken:
		return t.accessTokenService
	case RefreshToken:
		return t.refreshTokenService
	}

	return nil
}
