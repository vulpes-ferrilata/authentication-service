package handlers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/queries"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/results"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func NewGetRefreshTokenByClaimQueryHandler(tokenFactoryService services.TokenFactoryService) results.QueryHandlerResult {
	queryHandler := &getRefreshTokenByClaimQueryHandler{
		tokenFactoryService: tokenFactoryService,
	}

	return results.QueryHandlerResult{
		QueryHandler: queryHandler,
	}
}

type getRefreshTokenByClaimQueryHandler struct {
	tokenFactoryService services.TokenFactoryService
}

func (g getRefreshTokenByClaimQueryHandler) GetQuery() interface{} {
	return &queries.GetRefreshTokenByClaimQuery{}
}

func (g getRefreshTokenByClaimQueryHandler) Handle(ctx context.Context, query interface{}) (interface{}, error) {
	getRefreshTokenByClaimQuery := query.(*queries.GetRefreshTokenByClaimQuery)

	registeredClaim := &jwt.RegisteredClaims{
		ID:      getRefreshTokenByClaimQuery.JTI,
		Subject: getRefreshTokenByClaimQuery.UserID,
	}

	refreshToken, err := g.tokenFactoryService.GetTokenService(services.RefreshToken).Encrypt(ctx, registeredClaim)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return refreshToken, nil
}
