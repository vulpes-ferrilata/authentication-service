package queries

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/app_errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/services"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"github.com/vulpes-ferrilata/authentication-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTokenByClaimIDQuery struct {
	ClaimID string `validate:"required,objectid"`
}

func NewGetTokenByClaimIDQueryHandler(claimProjector projectors.ClaimProjector,
	tokenServiceResolver services.TokenServiceResolver) *GetTokenByClaimIDQueryHandler {
	return &GetTokenByClaimIDQueryHandler{
		claimProjector:       claimProjector,
		tokenServiceResolver: tokenServiceResolver,
	}
}

type GetTokenByClaimIDQueryHandler struct {
	claimProjector       projectors.ClaimProjector
	tokenServiceResolver services.TokenServiceResolver
}

func (g GetTokenByClaimIDQueryHandler) Handle(ctx context.Context, getTokenByClaimIDQuery *GetTokenByClaimIDQuery) (*models.Token, error) {
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
