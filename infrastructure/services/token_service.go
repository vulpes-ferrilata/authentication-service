package services

import (
	"context"
	"time"

	service "github.com/VulpesFerrilata/authentication-service"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type TokenService interface {
	Encrypt(ctx context.Context, registeredClaim *jwt.RegisteredClaims) (string, error)
	Decrypt(ctx context.Context, token string) (*jwt.RegisteredClaims, error)
}

func NewTokenService(signingMethod jwt.SigningMethod, secretKey string, duration time.Duration) TokenService {
	return &tokenService{
		signingMethod: signingMethod,
		secretKey:     secretKey,
		duration:      duration,
	}
}

type tokenService struct {
	signingMethod jwt.SigningMethod
	secretKey     string
	duration      time.Duration
}

func (t tokenService) Encrypt(ctx context.Context, registeredClaim *jwt.RegisteredClaims) (string, error) {
	registeredClaim.IssuedAt = jwt.NewNumericDate(time.Now())
	registeredClaim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(t.duration))

	token, err := jwt.NewWithClaims(t.signingMethod, registeredClaim).SignedString([]byte(t.secretKey))

	return token, errors.WithStack(err)
}

func (t tokenService) Decrypt(ctx context.Context, token string) (*jwt.RegisteredClaims, error) {
	registeredClaim := new(jwt.RegisteredClaims)

	parser := jwt.NewParser(
		jwt.WithoutClaimsValidation(),
		jwt.WithValidMethods([]string{t.signingMethod.Alg()}),
	)

	if _, err := parser.ParseWithClaims(token, registeredClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secretKey), nil
	}); err != nil {
		return nil, errors.WithStack(service.ErrTokenIsInvalid)
	}

	return registeredClaim, nil
}
