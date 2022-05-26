package grpc

import (
	"github.com/VulpesFerrilata/authentication-service/infrastructure/middlewares"
	"github.com/VulpesFerrilata/go-micro/plugins/server/grpc/v3"
	"github.com/VulpesFerrilata/shared/proto/authentication"
	"github.com/asim/go-micro/v3/server"
	"github.com/pkg/errors"
)

func NewServer(translatorMiddleware *middlewares.TranslatorMiddleware, errorHandlerMiddleware *middlewares.ErrorHandlerMiddleware, authenticationHandler authentication.AuthenticationHandler) (server.Server, error) {
	server := grpc.NewServer(
		server.WrapHandler(translatorMiddleware.WrapHandler),
		server.WrapHandler(errorHandlerMiddleware.WrapHandler),
	)

	if err := authentication.RegisterAuthenticationHandler(server, authenticationHandler); err != nil {
		return nil, errors.WithStack(err)
	}

	return server, nil
}
