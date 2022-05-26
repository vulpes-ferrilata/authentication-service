package services

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/domain/mappers"
	"github.com/VulpesFerrilata/authentication-service/domain/models"
	"github.com/VulpesFerrilata/authentication-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ClaimService interface {
	NewClaim(ctx context.Context, userID uuid.UUID, jti uuid.UUID) (*models.Claim, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Claim, error)
	Save(ctx context.Context, claim *models.Claim) error
	Delete(ctx context.Context, claim *models.Claim) error
}

func NewClaimService(claimRepository repositories.ClaimRepository,
	claimMapper mappers.ClaimMapper) ClaimService {
	return &claimService{
		claimRepository: claimRepository,
		claimMapper:     claimMapper,
	}
}

type claimService struct {
	claimRepository repositories.ClaimRepository
	claimMapper     mappers.ClaimMapper
}

func (c claimService) NewClaim(ctx context.Context, userID uuid.UUID, jti uuid.UUID) (*models.Claim, error) {
	claim := models.NewClaim(jti, userID)

	return claim, nil
}

func (c claimService) GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Claim, error) {
	claimEntity, err := c.claimRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	claim, err := c.claimMapper.ToModel(ctx, claimEntity)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return claim, nil
}

func (c claimService) Save(ctx context.Context, claim *models.Claim) error {
	claimEntity, err := c.claimMapper.ToEntity(ctx, claim)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := c.claimRepository.Save(ctx, claimEntity); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c claimService) Delete(ctx context.Context, claim *models.Claim) error {
	claimEntity, err := c.claimMapper.ToEntity(ctx, claim)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := c.claimRepository.Delete(ctx, claimEntity); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
