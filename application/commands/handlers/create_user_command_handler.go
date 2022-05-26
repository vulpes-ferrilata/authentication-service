package handlers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/commands"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/results"
	"github.com/VulpesFerrilata/shared/proto/user"
	"github.com/pkg/errors"
)

func NewCreateUserCommandHandler(userService user.UserService) results.CommandHandlerResult {
	commandHandler := &createUserCommandHandler{
		userService: userService,
	}

	return results.CommandHandlerResult{
		CommandHandler: commandHandler,
	}
}

type createUserCommandHandler struct {
	userService user.UserService
}

func (c createUserCommandHandler) GetCommand() interface{} {
	return &commands.CreateUserCommand{}
}

func (c createUserCommandHandler) Handle(ctx context.Context, command interface{}) error {
	createUserCommand := command.(*commands.CreateUserCommand)

	createUserRequest := &user.CreateUserRequest{
		ID:          createUserCommand.ID,
		DisplayName: createUserCommand.DisplayName,
	}

	if _, err := c.userService.CreateUser(ctx, createUserRequest); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
