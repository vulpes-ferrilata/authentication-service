package handlers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/queries"
	"github.com/VulpesFerrilata/authentication-service/application/queries/dtos"
	domain_services "github.com/VulpesFerrilata/authentication-service/domain/services"
	"github.com/VulpesFerrilata/authentication-service/domain/validators"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/results"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/services"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewGetClaimByRefreshTokenQueryHandler(claimValidator validators.ClaimValidator,
	claimService domain_services.ClaimService,
	tokenFactoryService services.TokenFactoryService) results.QueryHandlerResult {
	queryHandler := &getClaimByRefreshTokenQueryHandler{
		claimValidator:      claimValidator,
		claimService:        claimService,
		tokenFactoryService: tokenFactoryService,
	}

	return results.QueryHandlerResult{
		QueryHandler: queryHandler,
	}
}

type getClaimByRefreshTokenQueryHandler struct {
	claimValidator      validators.ClaimValidator
	claimService        domain_services.ClaimService
	tokenFactoryService services.TokenFactoryService
}

func (g getClaimByRefreshTokenQueryHandler) GetQuery() interface{} {
	return &queries.GetClaimByRefreshTokenQuery{}
}

func (g getClaimByRefreshTokenQueryHandler) Handle(ctx context.Context, query interface{}) (interface{}, error) {
	getClaimByRefreshTokenQuery := query.(*queries.GetClaimByRefreshTokenQuery)

	registeredClaim, err := g.tokenFactoryService.GetTokenService(services.RefreshToken).Decrypt(ctx, getClaimByRefreshTokenQuery.RefreshToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	jti, err := uuid.Parse(registeredClaim.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userID, err := uuid.Parse(registeredClaim.Subject)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := g.claimValidator.ValidateAuthentication(ctx, userID, jti); err != nil {
		return nil, errors.WithStack(err)
	}

	claim, err := g.claimService.GetByUserID(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	claimDTO := dtos.NewClaimDTO(claim)

	return claimDTO, nil
}
