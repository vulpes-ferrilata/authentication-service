package repositories

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/persistence/entities"
	"github.com/google/uuid"
)

type ClaimRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entities.Claim, error)
	Save(ctx context.Context, claimEntity *entities.Claim) error
	Delete(ctx context.Context, claimEntity *entities.Claim) error
}
