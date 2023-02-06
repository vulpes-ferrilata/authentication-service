package queries

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/app_errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/services"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"github.com/vulpes-ferrilata/authentication-service/view/projectors"
)

type GetClaimByAccessTokenQuery struct {
	AccessToken string `validate:"required,jwt"`
}

func NewGetClaimByAccessTokenQueryHandler(tokenServiceResolver services.TokenServiceResolver,
	claimProjector projectors.ClaimProjector) *GetClaimByAccessTokenQueryHandler {
	return &GetClaimByAccessTokenQueryHandler{
		tokenServiceResolver: tokenServiceResolver,
		claimProjector:       claimProjector,
	}
}

type GetClaimByAccessTokenQueryHandler struct {
	tokenServiceResolver services.TokenServiceResolver
	claimProjector       projectors.ClaimProjector
}

func (g GetClaimByAccessTokenQueryHandler) Handle(ctx context.Context, getClaimByAccessTokenQuery *GetClaimByAccessTokenQuery) (*models.Claim, error) {
	id, err := g.tokenServiceResolver.GetTokenService(services.AccessToken).Decrypt(ctx, getClaimByAccessTokenQuery.AccessToken)
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

	return claim, nil
}
