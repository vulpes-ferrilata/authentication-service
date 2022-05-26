package validators

import (
	"context"

	service "github.com/VulpesFerrilata/authentication-service"
	"github.com/VulpesFerrilata/authentication-service/persistence"
	"github.com/VulpesFerrilata/authentication-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ClaimValidator interface {
	ValidateAuthentication(ctx context.Context, userID uuid.UUID, jti uuid.UUID) error
}

func NewClaimValidator(claimRepository repositories.ClaimRepository) ClaimValidator {
	return &claimValidator{
		claimRepository: claimRepository,
	}
}

type claimValidator struct {
	claimRepository repositories.ClaimRepository
}

func (c claimValidator) ValidateAuthentication(ctx context.Context, userID uuid.UUID, jti uuid.UUID) error {
	claimEntity, err := c.claimRepository.GetByUserID(ctx, userID)
	if errors.Is(err, persistence.ErrRecordNotFound) {
		return errors.WithStack(service.ErrTokenHasBeenExpiredOrRevoked)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if claimEntity.JTI.String() != jti.String() {
		return errors.WithStack(service.ErrAccountHasBeenLoggedInByAnotherDevice)
	}

	return nil
}
