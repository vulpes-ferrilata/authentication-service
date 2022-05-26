package mappers

import (
	"context"
	"sync"

	"github.com/VulpesFerrilata/authentication-service/domain/models"
	"github.com/VulpesFerrilata/authentication-service/persistence/entities"
)

type ClaimMapper interface {
	ToEntity(ctx context.Context, claim *models.Claim) (*entities.Claim, error)
	ToModel(ctx context.Context, claimEntity *entities.Claim) (*models.Claim, error)
}

func NewClaimMapper() ClaimMapper {
	return &claimMapper{
		m: make(map[*models.Claim]*entities.Claim),
	}
}

type claimMapper struct {
	m  map[*models.Claim]*entities.Claim
	mu sync.RWMutex
}

func (c *claimMapper) ToEntity(ctx context.Context, claim *models.Claim) (*entities.Claim, error) {
	if claim == nil {
		return nil, nil
	}

	c.mu.RLock()
	claimEntity, ok := c.m[claim]
	c.mu.RUnlock()
	if !ok {
		claimEntity = new(entities.Claim)

		c.mu.Lock()
		c.m[claim] = claimEntity
		c.mu.Unlock()

		go func(claim *models.Claim, done <-chan struct{}) {
			<-done
			c.mu.Lock()
			delete(c.m, claim)
			c.mu.Unlock()
		}(claim, ctx.Done())
	}

	claimEntity.JTI = claim.GetJTI()
	claimEntity.UserID = claim.GetUserID()

	return claimEntity, nil
}

func (c *claimMapper) ToModel(ctx context.Context, claimEntity *entities.Claim) (*models.Claim, error) {
	if claimEntity == nil {
		return nil, nil
	}

	claim := models.NewClaim(
		claimEntity.JTI,
		claimEntity.UserID,
	)

	c.mu.Lock()
	c.m[claim] = claimEntity
	c.mu.Unlock()

	go func(claim *models.Claim, done <-chan struct{}) {
		<-done
		c.mu.Lock()
		delete(c.m, claim)
		c.mu.Unlock()
	}(claim, ctx.Done())

	return claim, nil
}
