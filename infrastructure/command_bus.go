package infrastructure

import (
	"github.com/vulpes-ferrilata/authentication-service/application/commands"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/middlewares"
	"github.com/vulpes-ferrilata/cqrs"
)

func NewCommandBus(validationMiddleware *middlewares.ValidationMiddleware,
	transactionMiddleware *middlewares.TransactionMiddleware,
	createUserCredentialCommandHandler *commands.CreateUserCredentialCommandHandler,
	deleteUserCredentialCommandHandler *commands.DeleteUserCredentialCommandHandler,
	loginCommandHandler *commands.LoginCommandHandler,
	revokeTokenCommandHandler *commands.RevokeTokenCommandHandler) (*cqrs.CommandBus, error) {
	commandBus := &cqrs.CommandBus{}

	commandBus.Use(
		validationMiddleware.CommandHandlerMiddleware(),
		transactionMiddleware.CommandHandlerMiddleware(),
	)

	commandBus.Register(&commands.CreateUserCredentialCommand{}, cqrs.WrapCommandHandlerFunc(createUserCredentialCommandHandler.Handle))
	commandBus.Register(&commands.DeleteUserCredentialCommand{}, cqrs.WrapCommandHandlerFunc(deleteUserCredentialCommandHandler.Handle))
	commandBus.Register(&commands.LoginCommand{}, cqrs.WrapCommandHandlerFunc(loginCommandHandler.Handle))
	commandBus.Register(&commands.RevokeTokenCommand{}, cqrs.WrapCommandHandlerFunc(revokeTokenCommandHandler.Handle))

	return commandBus, nil
}
