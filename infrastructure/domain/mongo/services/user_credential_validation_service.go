package services

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/domain/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserCredentialValidationService(db *mongo.Database) services.UserCredentialValidationService {
	return &userCredentialValidationService{
		userCredentialCollection: db.Collection("user_credentials"),
	}
}

type userCredentialValidationService struct {
	userCredentialCollection *mongo.Collection
}

func (u userCredentialValidationService) IsEmailAlreadyExists(ctx context.Context, email string) (bool, error) {
	filter := bson.M{"email": email}

	err := u.userCredentialCollection.FindOne(ctx, filter).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}
	if err != nil {
		return false, errors.WithStack(err)
	}

	return true, nil
}
