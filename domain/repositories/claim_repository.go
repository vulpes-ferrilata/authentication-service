package repositories

import (
	"context"

	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClaimRepository interface {
	Insert(ctx context.Context, claim *models.Claim) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
