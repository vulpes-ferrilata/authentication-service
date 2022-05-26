package grpc

import (
	"github.com/VulpesFerrilata/authentication-service/infrastructure"
	grpc_handlers "github.com/VulpesFerrilata/authentication-service/presentation/grpc/handlers"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := infrastructure.NewContainer()

	//Presentation layer
	container.Provide(grpc_handlers.NewAuthenticationHandler)

	container.Provide(NewServer)
	container.Provide(NewClient)
	container.Provide(NewService)

	return container
}
