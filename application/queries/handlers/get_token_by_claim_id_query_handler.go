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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewGetTokenByClaimIDQueryHandler(validate *validator.Validate,
	claimProjector projectors.ClaimProjector,
	tokenServiceResolver services.TokenServiceResolver) query.QueryHandler[*queries.GetTokenByClaimIDQuery, *models.Token] {
	handler := &getTokenByClaimIDQueryHandler{
		claimProjector:       claimProjector,
		tokenServiceResolver: tokenServiceResolver,
	}
	validationWrapper := wrappers.NewValidationWrapper[*queries.GetTokenByClaimIDQuery, *models.Token](validate, handler)

	return validationWrapper
}

type getTokenByClaimIDQueryHandler struct {
	claimProjector       projectors.ClaimProjector
	tokenServiceResolver services.TokenServiceResolver
}

func (g getTokenByClaimIDQueryHandler) Handle(ctx context.Context, getTokenByClaimIDQuery *queries.GetTokenByClaimIDQuery) (*models.Token, error) {
	id, err := primitive.ObjectIDFromHex(getTokenByClaimIDQuery.ClaimID)
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

	refreshToken, err := g.tokenServiceResolver.GetTokenService(services.RefreshToken).Encrypt(ctx, claim.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	token := &models.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
