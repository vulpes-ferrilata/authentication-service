package rest

import (
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
)

func NewApp(router Router) (*iris.Application, error) {
	app := iris.Default()
	router.Init(app)

	if err := app.Build(); err != nil {
		return nil, errors.WithStack(err)
	}

	return app, nil
}
