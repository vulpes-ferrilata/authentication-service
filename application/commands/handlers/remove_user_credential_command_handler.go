package handlers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/commands"
	domain_services "github.com/VulpesFerrilata/authentication-service/domain/services"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/results"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewRemoveUserCredentialCommandHandler(userCredentialService domain_services.UserCredentialService) results.CommandHandlerResult {
	commandHandler := &removeUserCredentialCommandHandler{
		userCredentialService: userCredentialService,
	}

	return results.CommandHandlerResult{
		CommandHandler: commandHandler,
	}
}

type removeUserCredentialCommandHandler struct {
	userCredentialService domain_services.UserCredentialService
}

func (r removeUserCredentialCommandHandler) GetCommand() interface{} {
	return &commands.RemoveUserCredentialCommand{}
}

func (r removeUserCredentialCommandHandler) Handle(ctx context.Context, command interface{}) error {
	removeUserCredentialCommand := command.(*commands.RemoveUserCredentialCommand)

	id, err := uuid.Parse(removeUserCredentialCommand.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	userCredential, err := r.userCredentialService.GetByID(ctx, id)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := r.userCredentialService.Delete(ctx, userCredential); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
