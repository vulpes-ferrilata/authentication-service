package repositories

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/domain/mongo/documents"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/domain/mongo/mappers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserCredentialRepository(db *mongo.Database) repositories.UserCredentialRepository {
	return &userCredentialRepository{
		userCredentialCollection: db.Collection("user_credentials"),
	}
}

type userCredentialRepository struct {
	userCredentialCollection *mongo.Collection
}

func (u userCredentialRepository) GetByEmail(ctx context.Context, email string) (*models.UserCredential, error) {
	userCredentialDocument := &documents.UserCredential{}

	filter := bson.M{"email": email}

	err := u.userCredentialCollection.FindOne(ctx, filter).Decode(userCredentialDocument)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, errors.WithStack(app_errors.ErrUserCredentialNotFound)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userCredential := mappers.ToUserCredentialDomain(userCredentialDocument)

	return userCredential, nil
}

func (u userCredentialRepository) Insert(ctx context.Context, userCredential *models.UserCredential) error {
	userCredentialDocument := mappers.ToUserCredentialDocument(userCredential)

	userCredentialDocument.Version = 1

	if _, err := u.userCredentialCollection.InsertOne(ctx, userCredentialDocument); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (u userCredentialRepository) Update(ctx context.Context, userCredential *models.UserCredential) error {
	userCredentialDocument := mappers.ToUserCredentialDocument(userCredential)

	filter := bson.M{"_id": userCredentialDocument.ID, "version": userCredentialDocument.Version}

	userCredentialDocument.Version++

	result, err := u.userCredentialCollection.ReplaceOne(ctx, filter, userCredentialDocument)
	if err != nil {
		return errors.WithStack(err)
	}
	if result.ModifiedCount == 0 {
		return errors.WithStack(app_errors.ErrStaleUserCredential)
	}

	return nil
}

func (u userCredentialRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	result, err := u.userCredentialCollection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.WithStack(err)
	}
	if result.DeletedCount == 0 {
		return errors.WithStack(app_errors.ErrStaleUserCredential)
	}

	return nil
}
