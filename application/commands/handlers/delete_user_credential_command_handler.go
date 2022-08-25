package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/application/commands"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command/wrappers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewDeleteUserCredentialCommandHandler(validate *validator.Validate, db *mongo.Database, userCredentialRepository repositories.UserCredentialRepository) command.CommandHandler[*commands.DeleteUserCredentialCommand] {
	handler := &deleteUserCredentialCommandHandler{
		userCredentialRepository: userCredentialRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.DeleteUserCredentialCommand](db, handler)
	validationWrapper := wrappers.NewValidationWrapper[*commands.DeleteUserCredentialCommand](validate, transactionWrapper)

	return validationWrapper
}

type deleteUserCredentialCommandHandler struct {
	userCredentialRepository repositories.UserCredentialRepository
}

func (d deleteUserCredentialCommandHandler) Handle(ctx context.Context, deleteUserCredentialCommand *commands.DeleteUserCredentialCommand) error {
	id, err := primitive.ObjectIDFromHex(deleteUserCredentialCommand.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := d.userCredentialRepository.Delete(ctx, id); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
