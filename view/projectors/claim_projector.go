package projectors

import (
	"context"

	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClaimProjector interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.Claim, error)
}
