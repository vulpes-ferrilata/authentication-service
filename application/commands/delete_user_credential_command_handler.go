package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteUserCredentialCommand struct {
	UserCredentialID string `validate:"required,objectid"`
}

func NewDeleteUserCredentialCommandHandler(userCredentialRepository repositories.UserCredentialRepository) *DeleteUserCredentialCommandHandler {
	return &DeleteUserCredentialCommandHandler{
		userCredentialRepository: userCredentialRepository,
	}
}

type DeleteUserCredentialCommandHandler struct {
	userCredentialRepository repositories.UserCredentialRepository
}

func (d DeleteUserCredentialCommandHandler) Handle(ctx context.Context, deleteUserCredentialCommand *DeleteUserCredentialCommand) error {
	userCredentialID, err := primitive.ObjectIDFromHex(deleteUserCredentialCommand.UserCredentialID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := d.userCredentialRepository.Delete(ctx, userCredentialID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
