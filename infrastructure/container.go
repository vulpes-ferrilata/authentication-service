package infrastructure

import (
	command_handlers "github.com/VulpesFerrilata/authentication-service/application/commands/handlers"
	query_handlers "github.com/VulpesFerrilata/authentication-service/application/queries/handlers"
	"github.com/VulpesFerrilata/authentication-service/domain/mappers"
	domain_services "github.com/VulpesFerrilata/authentication-service/domain/services"
	"github.com/VulpesFerrilata/authentication-service/domain/validators"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/middlewares"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/persistence/repositories"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/services"
	persistence_repositories "github.com/VulpesFerrilata/authentication-service/persistence/repositories"
	"github.com/VulpesFerrilata/shared/proto/user"
	"github.com/asim/go-micro/v3/client"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//Other services
	container.Provide(func(config *Config) services.TokenFactoryService {
		accessTokenService := services.NewTokenService(jwt.SigningMethodHS512, config.Authentication.AccessTokenSecret, config.Authentication.AccessTokenDuration)
		refreshTokenService := services.NewTokenService(jwt.SigningMethodHS512, config.Authentication.RefreshTokenSecret, config.Authentication.RefreshTokenDuration)
		return services.NewTokenFactoryService(accessTokenService, refreshTokenService)
	})

	//3rd party libraries
	container.Provide(NewConfig)
	container.Provide(NewRedis)
	container.Provide(NewGorm)
	container.Provide(NewValidate)
	container.Provide(NewCommandBus)
	container.Provide(NewQueryBus)
	container.Provide(NewUniversalTranslator)

	//middlewares
	container.Provide(middlewares.NewErrorHandlerMiddleware)
	container.Provide(middlewares.NewValidationMiddleware)
	container.Provide(middlewares.NewTransactionMiddleware)
	container.Provide(middlewares.NewTranslatorMiddleware)

	//Persistence layer
	container.Provide(repositories.NewUserCredentialRepository)
	container.Provide(func(rdb *redis.Client, config *Config) persistence_repositories.ClaimRepository {
		return repositories.NewClaimRepository(rdb, config.Authentication.RefreshTokenDuration)
	})

	//Domain layer
	//--Services
	container.Provide(domain_services.NewUserCredentialService)
	container.Provide(domain_services.NewClaimService)
	container.Provide(domain_services.NewUserService)
	//--Mappers
	container.Provide(mappers.NewUserCredentialMapper)
	container.Provide(mappers.NewClaimMapper)
	//--Validators
	container.Provide(validators.NewUserCredentialValidator)
	container.Provide(validators.NewClaimValidator)

	//Micro services
	container.Provide(func(client client.Client) user.UserService {
		return user.NewUserService("boardgame.user.service", client)
	})

	//Application layer
	//--Queries
	container.Provide(query_handlers.NewGetAccessTokenByClaimQueryHandler)
	container.Provide(query_handlers.NewGetClaimByAccessTokenQueryHandler)
	container.Provide(query_handlers.NewGetRefreshTokenByClaimQueryHandler)
	container.Provide(query_handlers.NewGetClaimByRefreshTokenQueryHandler)
	container.Provide(query_handlers.NewGetUserCredentialByEmailAndPasswordQueryHandler)
	//--Commands
	container.Provide(command_handlers.NewCreateClaimCommandHandler)
	container.Provide(command_handlers.NewCreateUserCredentialCommandHandler)
	container.Provide(command_handlers.NewCreateUserCommandHandler)
	container.Provide(command_handlers.NewRemoveClaimCommandHandler)
	container.Provide(command_handlers.NewRemoveUserCredentialCommandHandler)

	return container
}
