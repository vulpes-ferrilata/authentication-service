package repositories

import (
	"context"

	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserCredentialRepository interface {
	GetByEmail(ctx context.Context, email string) (*models.UserCredential, error)
	Insert(ctx context.Context, userCredential *models.UserCredential) error
	Update(ctx context.Context, userCredential *models.UserCredential) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
