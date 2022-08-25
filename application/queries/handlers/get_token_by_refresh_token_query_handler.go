package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/application/queries"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/query/wrappers"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/services"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"github.com/vulpes-ferrilata/authentication-service/view/projectors"
)

func NewGetTokenByRefreshTokenQueryHandler(validate *validator.Validate, tokenServiceResolver services.TokenServiceResolver,
	claimProjector projectors.ClaimProjector) query.QueryHandler[*queries.GetTokenByRefreshTokenQuery, *models.Token] {
	handler := &getTokenByRefreshTokenQueryHandler{
		tokenServiceResolver: tokenServiceResolver,
		claimProjector:       claimProjector,
	}
	validationWrapper := wrappers.NewValidationWrapper[*queries.GetTokenByRefreshTokenQuery, *models.Token](validate, handler)

	return validationWrapper
}

type getTokenByRefreshTokenQueryHandler struct {
	tokenServiceResolver services.TokenServiceResolver
	claimProjector       projectors.ClaimProjector
}

func (g getTokenByRefreshTokenQueryHandler) Handle(ctx context.Context, getTokenByRefreshTokenQuery *queries.GetTokenByRefreshTokenQuery) (*models.Token, error) {
	id, err := g.tokenServiceResolver.GetTokenService(services.RefreshToken).Decrypt(ctx, getTokenByRefreshTokenQuery.RefreshToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	claim, err := g.claimProjector.GetByID(ctx, id)
	if errors.Is(err, app_errors.ErrClaimNotFound) {
		return nil, errors.WithStack(app_errors.ErrTokenHasBeenRevoked)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	accessToken, err := g.tokenServiceResolver.GetTokenService(services.AccessToken).Encrypt(ctx, claim.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	token := &models.Token{
		AccessToken: accessToken,
	}

	return token, nil
}
