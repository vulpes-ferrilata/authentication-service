package rest

import (
	"github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
)

func NewServer(app *iris.Application) (server.Server, error) {
	server := http.NewServer()

	if err := micro.RegisterHandler(server, app); err != nil {
		return nil, errors.WithStack(err)
	}

	return server, nil
}
