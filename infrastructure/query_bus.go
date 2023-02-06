package infrastructure

import (
	"github.com/vulpes-ferrilata/authentication-service/application/queries"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/middlewares"
	"github.com/vulpes-ferrilata/cqrs"
)

func NewQueryBus(validationMiddleware *middlewares.ValidationMiddleware,
	getClaimByAccessTokenQueryHandler *queries.GetClaimByAccessTokenQueryHandler,
	getTokenByClaimIDQueryHandler *queries.GetTokenByClaimIDQueryHandler,
	getTokenByRefreshTokenQueryHandler *queries.GetTokenByRefreshTokenQueryHandler) (*cqrs.QueryBus, error) {
	queryBus := &cqrs.QueryBus{}

	queryBus.Use(
		validationMiddleware.QueryHandlerMiddleware(),
	)

	queryBus.Register(&queries.GetClaimByAccessTokenQuery{}, cqrs.WrapQueryHandlerFunc(getClaimByAccessTokenQueryHandler.Handle))
	queryBus.Register(&queries.GetTokenByClaimIDQuery{}, cqrs.WrapQueryHandlerFunc(getTokenByClaimIDQueryHandler.Handle))
	queryBus.Register(&queries.GetTokenByRefreshTokenQuery{}, cqrs.WrapQueryHandlerFunc(getTokenByRefreshTokenQueryHandler.Handle))

	return queryBus, nil
}
