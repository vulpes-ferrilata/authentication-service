package presentation

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/shared/proto/v1/authentication"
	"google.golang.org/grpc"
)

func NewServer(logger *logrus.Logger,
	recoverInterceptor *interceptors.RecoverInterceptor,
	errorHandlerInterceptor *interceptors.ErrorHandlerInterceptor,
	localeInterceptor *interceptors.LocaleInterceptor,
	authenticationServer authentication.AuthenticationServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
			recoverInterceptor.ServerUnaryInterceptor(),
			errorHandlerInterceptor.ServerUnaryInterceptor(),
			localeInterceptor.ServerUnaryInterceptor(),
		),
	)

	authentication.RegisterAuthenticationServer(server, authenticationServer)

	return server
}
