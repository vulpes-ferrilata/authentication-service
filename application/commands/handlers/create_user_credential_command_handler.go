package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/application/commands"
	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command/wrappers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCreateUserCredentialCommandHandler(validate *validator.Validate, db *mongo.Database, userCredentialRepository repositories.UserCredentialRepository) command.CommandHandler[*commands.CreateUserCredentialCommand] {
	handler := &createUserCredentialCommandHandler{
		userCredentialRepository: userCredentialRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.CreateUserCredentialCommand](db, handler)
	validationWrapper := wrappers.NewValidationWrapper[*commands.CreateUserCredentialCommand](validate, transactionWrapper)

	return validationWrapper
}

type createUserCredentialCommandHandler struct {
	userCredentialRepository repositories.UserCredentialRepository
}

func (c createUserCredentialCommandHandler) Handle(ctx context.Context, createUserCredentialCommand *commands.CreateUserCredentialCommand) error {
	id, err := primitive.ObjectIDFromHex(createUserCredentialCommand.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(createUserCredentialCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	userCredential := models.NewUserCredentialBuilder().
		SetID(id).
		SetUserID(userID).
		SetEmail(createUserCredentialCommand.Email).
		Create()

	if err := userCredential.SetPassword(createUserCredentialCommand.Password); err != nil {
		return errors.WithStack(err)
	}

	if err := c.userCredentialRepository.Insert(ctx, userCredential); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
