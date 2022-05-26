package projectors

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/persistence/repositories"
	"github.com/VulpesFerrilata/authentication-service/view/models"
	"github.com/google/uuid"
)

type ClaimProjector interface {
	GetByID(ctx context.Context, userID uuid.UUID) (*models.Claim, error)
}

func NewClaimProjector() ClaimProjector {
	return &claimProjector{}
}

type claimProjector struct {
	claimRepository repositories.ClaimRepository
}

func (c claimProjector) GetByID(ctx context.Context, userID uuid.UUID) (*models.Claim, error) {

}
