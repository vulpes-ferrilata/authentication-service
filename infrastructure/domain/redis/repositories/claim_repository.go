package repositories

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/config"
	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	DuplicateKeyErr = errors.New("duplicate key")
)

func NewClaimRepository(redisClient *redis.Client, config config.Config) (repositories.ClaimRepository, error) {
	expirationDuration, err := time.ParseDuration(config.RefreshToken.Expiration)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &claimRepository{
		redisClient:        redisClient,
		expirationDuration: expirationDuration,
	}, nil
}

type claimRepository struct {
	redisClient        *redis.Client
	expirationDuration time.Duration
}

func (c claimRepository) Insert(ctx context.Context, claim *models.Claim) error {
	ok, err := c.redisClient.SetNX(ctx, claim.GetID().Hex(), claim.GetUserID().Hex(), c.expirationDuration).Result()
	if err != nil {
		return errors.WithStack(err)
	}
	if !ok {
		return errors.WithStack(DuplicateKeyErr)
	}

	return nil
}

func (c claimRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	if err := c.redisClient.Del(ctx, id.Hex()).Err(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
