package infrastructure

import (
	command_handlers "github.com/vulpes-ferrilata/authentication-service/application/commands/handlers"
	query_handlers "github.com/vulpes-ferrilata/authentication-service/application/queries/handlers"
	mongo_repositories "github.com/vulpes-ferrilata/authentication-service/infrastructure/domain/mongo/repositories"
	mongo_services "github.com/vulpes-ferrilata/authentication-service/infrastructure/domain/mongo/services"
	redis_repositories "github.com/vulpes-ferrilata/authentication-service/infrastructure/domain/redis/repositories"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/services"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/view/redis/projectors"
	"github.com/vulpes-ferrilata/authentication-service/presentation/v1/servers"
	"go.uber.org/dig"
	"google.golang.org/grpc"
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
	container.Provide(grpc.NewServer)
	//--3rd party services
	container.Provide(services.NewTokenServiceResolver)
	//--Grpc interceptors
	container.Provide(interceptors.NewRecoverInterceptor)
	container.Provide(interceptors.NewErrorHandlerInterceptor)
	container.Provide(interceptors.NewLocaleInterceptor)

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
	container.Provide(query_handlers.NewGetClaimByAccessTokenQueryHandler)
	container.Provide(query_handlers.NewGetTokenByClaimIDQueryHandler)
	container.Provide(query_handlers.NewGetTokenByRefreshTokenQueryHandler)
	//--Commands
	container.Provide(command_handlers.NewCreateUserCredentialCommandHandler)
	container.Provide(command_handlers.NewLoginCommandHandler)
	container.Provide(command_handlers.NewDeleteUserCredentialCommandHandler)
	container.Provide(command_handlers.NewRevokeTokenCommandHandler)

	//Presentation layer
	container.Provide(servers.NewAuthenticationServer)

	return container
}
