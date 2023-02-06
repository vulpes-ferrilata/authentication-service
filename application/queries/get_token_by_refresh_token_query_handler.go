package queries

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/app_errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/services"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"github.com/vulpes-ferrilata/authentication-service/view/projectors"
)

type GetTokenByRefreshTokenQuery struct {
	RefreshToken string `validate:"required,jwt"`
}

func NewGetTokenByRefreshTokenQueryHandler(tokenServiceResolver services.TokenServiceResolver,
	claimProjector projectors.ClaimProjector) *GetTokenByRefreshTokenQueryHandler {
	return &GetTokenByRefreshTokenQueryHandler{
		tokenServiceResolver: tokenServiceResolver,
		claimProjector:       claimProjector,
	}
}

type GetTokenByRefreshTokenQueryHandler struct {
	tokenServiceResolver services.TokenServiceResolver
	claimProjector       projectors.ClaimProjector
}

func (g GetTokenByRefreshTokenQueryHandler) Handle(ctx context.Context, getTokenByRefreshTokenQuery *GetTokenByRefreshTokenQuery) (*models.Token, error) {
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
