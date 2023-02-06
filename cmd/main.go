package main

import (
	"net"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/config"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure"
	"google.golang.org/grpc"
)

func main() {
	container := infrastructure.NewContainer()

	if err := container.Invoke(func(server *grpc.Server, config config.Config) error {
		listener, err := net.Listen("tcp", config.Server.Address)
		if err != nil {
			return errors.WithStack(err)
		}

		if err := server.Serve(listener); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		panic(err)
	}
}
