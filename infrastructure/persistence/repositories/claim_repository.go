package repositories

import (
	"context"
	"time"

	"github.com/VulpesFerrilata/authentication-service/persistence"
	"github.com/VulpesFerrilata/authentication-service/persistence/entities"
	"github.com/VulpesFerrilata/authentication-service/persistence/repositories"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewClaimRepository(rdb *redis.Client, duration time.Duration) repositories.ClaimRepository {
	return &claimRepository{
		rdb:      rdb,
		duration: duration,
	}
}

type claimRepository struct {
	rdb      *redis.Client
	duration time.Duration
}

func (c claimRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entities.Claim, error) {
	claim := &entities.Claim{
		UserID: userID,
	}

	result, err := c.rdb.Get(ctx, userID.String()).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errors.WithStack(persistence.ErrRecordNotFound)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	jti, err := uuid.Parse(result)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	claim.JTI = jti

	return claim, nil
}

func (c claimRepository) Save(ctx context.Context, claimEntity *entities.Claim) error {
	if err := c.rdb.Watch(ctx, func(t *redis.Tx) error {
		if _, err := t.TxPipelined(ctx, func(p redis.Pipeliner) error {
			if err := p.Set(ctx, claimEntity.UserID.String(), claimEntity.JTI.String(), c.duration).Err(); err != nil {
				return errors.WithStack(err)
			}

			return nil
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}, claimEntity.UserID.String()); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c claimRepository) Delete(ctx context.Context, claimEntity *entities.Claim) error {
	if err := c.rdb.Watch(ctx, func(t *redis.Tx) error {
		result, err := t.Get(ctx, claimEntity.UserID.String()).Result()
		if errors.Is(err, redis.Nil) {
			return nil
		}
		if err != nil {
			return errors.WithStack(err)
		}

		if result != claimEntity.JTI.String() {
			return nil
		}

		if _, err = t.TxPipelined(ctx, func(p redis.Pipeliner) error {
			if err := p.Del(ctx, claimEntity.UserID.String()).Err(); err != nil {
				return errors.WithStack(err)
			}

			return nil
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}, claimEntity.UserID.String()); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
