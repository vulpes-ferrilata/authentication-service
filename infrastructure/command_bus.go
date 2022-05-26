package infrastructure

import (
	"github.com/VulpesFerrilata/authentication-service/infrastructure/bus"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/params"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/middlewares"
)

func NewCommandBus(params params.CommandBusParams,
	validationMiddleware *middlewares.ValidationMiddleware,
	transactionMiddleware *middlewares.TransactionMiddleware) bus.CommandBus {
	commandBus := bus.NewCommandBus()

	commandBus.Register(params.CommandHandlers...)
	commandBus.Use(
		transactionMiddleware.WrapCommandHandler,
		validationMiddleware.WrapCommandHandler,
	)

	return commandBus
}
