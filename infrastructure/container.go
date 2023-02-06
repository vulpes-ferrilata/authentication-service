package infrastructure

import (
	"github.com/vulpes-ferrilata/authentication-service/application/commands"
	"github.com/vulpes-ferrilata/authentication-service/application/queries"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/middlewares"
	mongo_repositories "github.com/vulpes-ferrilata/authentication-service/infrastructure/domain/mongo/repositories"
	mongo_services "github.com/vulpes-ferrilata/authentication-service/infrastructure/domain/mongo/services"
	redis_repositories "github.com/vulpes-ferrilata/authentication-service/infrastructure/domain/redis/repositories"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/services"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/view/redis/projectors"
	"github.com/vulpes-ferrilata/authentication-service/presentation"
	v1 "github.com/vulpes-ferrilata/authentication-service/presentation/v1"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//Infrastructure layer
	container.Provide(NewConfig)
	container.Provide(NewRedis)
	container.Provide(NewMongo)
	container.Provide(NewValidator)
	container.Provide(NewLogrus)
	container.Provide(NewUniversalTranslator)

	//--3rd party services
	container.Provide(services.NewTokenServiceResolver)
	//--Grpc interceptors
	container.Provide(interceptors.NewRecoverInterceptor)
	container.Provide(interceptors.NewErrorHandlerInterceptor)
	container.Provide(interceptors.NewLocaleInterceptor)
	//--Cqrs middlewares
	container.Provide(middlewares.NewValidationMiddleware)
	container.Provide(middlewares.NewTransactionMiddleware)

	//Domain layer
	//--Repositories
	container.Provide(mongo_repositories.NewUserCredentialRepository)
	container.Provide(redis_repositories.NewClaimRepository)
	//--Services
	container.Provide(mongo_services.NewUserCredentialValidationService)

	//View layer
	//--Projectors
	container.Provide(projectors.NewClaimProjector)

	//Application layer
	//--Queries
	container.Provide(queries.NewGetClaimByAccessTokenQueryHandler)
	container.Provide(queries.NewGetTokenByClaimIDQueryHandler)
	container.Provide(queries.NewGetTokenByRefreshTokenQueryHandler)
	//--Commands
	container.Provide(commands.NewCreateUserCredentialCommandHandler)
	container.Provide(commands.NewLoginCommandHandler)
	container.Provide(commands.NewDeleteUserCredentialCommandHandler)
	container.Provide(commands.NewRevokeTokenCommandHandler)

	//Presentation layer
	container.Provide(presentation.NewServer)
	container.Provide(v1.NewAuthenticationServer)

	return container
}
