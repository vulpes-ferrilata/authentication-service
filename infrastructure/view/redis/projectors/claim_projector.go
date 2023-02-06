package projectors

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/app_errors"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"github.com/vulpes-ferrilata/authentication-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewClaimProjector(rdb *redis.Client) projectors.ClaimProjector {
	return &claimProjector{
		rdb: rdb,
	}
}

type claimProjector struct {
	rdb *redis.Client
}

func (c claimProjector) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Claim, error) {
	userID, err := c.rdb.Get(ctx, id.Hex()).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errors.WithStack(app_errors.ErrClaimNotFound)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	claim := &models.Claim{
		ID:     id,
		UserID: userObjectID,
	}

	return claim, nil
}
