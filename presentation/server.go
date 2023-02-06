package presentation

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"github.com/vulpes-ferrilata/authentication-service-proto/pb"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/grpc/interceptors"
	"google.golang.org/grpc"
)

func NewServer(logger *logrus.Logger,
	recoverInterceptor *interceptors.RecoverInterceptor,
	errorHandlerInterceptor *interceptors.ErrorHandlerInterceptor,
	localeInterceptor *interceptors.LocaleInterceptor,
	authenticationServer pb.AuthenticationServer) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
			recoverInterceptor.ServerUnaryInterceptor(),
			localeInterceptor.ServerUnaryInterceptor(),
			errorHandlerInterceptor.ServerUnaryInterceptor(),
		),
	)

	pb.RegisterAuthenticationServer(server, authenticationServer)

	return server
}
