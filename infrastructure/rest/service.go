package rest

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/server"
)

func NewService(server server.Server, client client.Client) micro.Service {
	service := micro.NewService(
		micro.Server(
			server,
		),
		micro.Client(
			client,
		),
		micro.Name("boardgame.authentication.web"),
		micro.Version("latest"),
	)

	service.Init()

	return service
}
