package main

import (
	"github.com/VulpesFerrilata/authentication-service/infrastructure/rest"
	"github.com/asim/go-micro/v3"
	"github.com/pkg/errors"
)

func main() {
	container := rest.NewContainer()

	if err := container.Invoke(func(service micro.Service) error {
		if err := service.Run(); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		panic(err)
	}
}
