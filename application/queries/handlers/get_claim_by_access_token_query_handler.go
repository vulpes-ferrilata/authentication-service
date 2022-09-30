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

func NewGetClaimByAccessTokenQueryHandler(validate *validator.Validate,
	tokenServiceResolver services.TokenServiceResolver,
	claimProjector projectors.ClaimProjector) query.QueryHandler[*queries.GetClaimByAccessToken, *models.Claim] {
	handler := &getClaimByAccessTokenQueryHandler{
		tokenServiceResolver: tokenServiceResolver,
		claimProjector:       claimProjector,
	}
	validationWrapper := wrappers.NewValidationWrapper[*queries.GetClaimByAccessToken, *models.Claim](validate, handler)

	return validationWrapper
}

type getClaimByAccessTokenQueryHandler struct {
	tokenServiceResolver services.TokenServiceResolver
	claimProjector       projectors.ClaimProjector
}

func (g getClaimByAccessTokenQueryHandler) Handle(ctx context.Context, getClaimByAccessTokenQuery *queries.GetClaimByAccessToken) (*models.Claim, error) {
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
