package handlers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/queries"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/results"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func NewGetAccessTokenByClaimQueryHandler(tokenFactoryService services.TokenFactoryService) results.QueryHandlerResult {
	queryHandler := &getAccessTokenByClaimQueryHandler{
		tokenFactoryService: tokenFactoryService,
	}

	return results.QueryHandlerResult{
		QueryHandler: queryHandler,
	}
}

type getAccessTokenByClaimQueryHandler struct {
	tokenFactoryService services.TokenFactoryService
}

func (g getAccessTokenByClaimQueryHandler) GetQuery() interface{} {
	return &queries.GetAccessTokenByClaimQuery{}
}

func (g getAccessTokenByClaimQueryHandler) Handle(ctx context.Context, query interface{}) (interface{}, error) {
	getAccessTokenByClaimQuery := query.(*queries.GetAccessTokenByClaimQuery)

	registeredClaim := &jwt.RegisteredClaims{
		ID:      getAccessTokenByClaimQuery.JTI,
		Subject: getAccessTokenByClaimQuery.UserID,
	}

	accessToken, err := g.tokenFactoryService.GetTokenService(services.AccessToken).Encrypt(ctx, registeredClaim)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return accessToken, nil
}
