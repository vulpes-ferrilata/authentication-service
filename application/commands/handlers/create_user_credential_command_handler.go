package handlers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/commands"
	"github.com/VulpesFerrilata/authentication-service/domain/services"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/results"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewCreateUserCredentialCommandHandler(validate *validator.Validate,
	userCredentialService services.UserCredentialService,
	userService services.UserService) results.CommandHandlerResult {
	commandHandler := &createUserCredentialCommandHandler{
		userCredentialService: userCredentialService,
		userService:           userService,
	}

	return results.CommandHandlerResult{
		CommandHandler: commandHandler,
	}
}

type createUserCredentialCommandHandler struct {
	userCredentialService services.UserCredentialService
	userService           services.UserService
}

func (c createUserCredentialCommandHandler) GetCommand() interface{} {
	return &commands.CreateUserCredentialCommand{}
}

func (c createUserCredentialCommandHandler) Handle(ctx context.Context, command interface{}) error {
	createUserCredentialCommand := command.(*commands.CreateUserCredentialCommand)

	id, err := uuid.Parse(createUserCredentialCommand.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := uuid.Parse(createUserCredentialCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	userCredential, err := c.userCredentialService.NewUserCredential(ctx, id, userID, createUserCredentialCommand.Email, createUserCredentialCommand.Password)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := c.userCredentialService.Save(ctx, userCredential); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
