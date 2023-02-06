package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/app_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenService interface {
	Encrypt(ctx context.Context, id primitive.ObjectID) (string, error)
	Decrypt(ctx context.Context, token string) (primitive.ObjectID, error)
}

func NewTokenService(algorithm string, secretKey string, expiration string) (TokenService, error) {
	signingMethod := jwt.GetSigningMethod(algorithm)
	if signingMethod == nil {
		return nil, errors.Errorf("%s is invalid signing method", algorithm)
	}

	expirationDuration, err := time.ParseDuration(expiration)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &tokenService{
		signingMethod:      signingMethod,
		secretKey:          secretKey,
		expirationDuration: expirationDuration,
	}, nil
}

type tokenService struct {
	signingMethod      jwt.SigningMethod
	secretKey          string
	expirationDuration time.Duration
}

func (t tokenService) Encrypt(ctx context.Context, id primitive.ObjectID) (string, error) {
	registeredClaim := &jwt.RegisteredClaims{
		ID:        id.Hex(),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.expirationDuration)),
	}

	token, err := jwt.NewWithClaims(t.signingMethod, registeredClaim).SignedString([]byte(t.secretKey))

	return token, errors.WithStack(err)
}

func (t tokenService) Decrypt(ctx context.Context, token string) (primitive.ObjectID, error) {
	registeredClaim := new(jwt.RegisteredClaims)

	parser := jwt.NewParser(
		jwt.WithoutClaimsValidation(),
		jwt.WithValidMethods([]string{t.signingMethod.Alg()}),
	)

	if _, err := parser.ParseWithClaims(token, registeredClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secretKey), nil
	}); err != nil {
		return primitive.NilObjectID, app_errors.ErrTokenIsInvalid
	}

	if !registeredClaim.VerifyExpiresAt(time.Now(), true) {
		return primitive.NilObjectID, app_errors.ErrTokenHasBeenExpired
	}

	id, err := primitive.ObjectIDFromHex(registeredClaim.ID)
	if err != nil {
		return primitive.NilObjectID, errors.WithStack(err)
	}

	return id, nil
}
