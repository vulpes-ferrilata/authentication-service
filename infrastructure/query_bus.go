package infrastructure

import (
	"github.com/VulpesFerrilata/authentication-service/infrastructure/bus"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/params"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/middlewares"
)

func NewQueryBus(params params.QueryBusParams,
	validationMiddleware *middlewares.ValidationMiddleware,
	transactionMiddleware *middlewares.TransactionMiddleware) bus.QueryBus {
	queryBus := bus.NewQueryBus()

	queryBus.Register(params.QueryHandlers...)
	queryBus.Use(
		transactionMiddleware.WrapQueryHandler,
		validationMiddleware.WrapQueryHandler,
	)

	return queryBus
}
