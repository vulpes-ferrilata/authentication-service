package middlewares

import (
	"context"
	"fmt"

	micro_errors "github.com/asim/go-micro/v3/errors"
	"github.com/asim/go-micro/v3/server"
	"github.com/kataras/iris/v12"
)

func NewRecoverMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{}
}

type RecoverMiddleware struct{}

func (r RecoverMiddleware) WrapHandler(f server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = micro_errors.InternalServerError(req.Endpoint(), "%+v", r)
			}
		}()

		err = f(ctx, req, rsp)

		return
	}
}

func (r RecoverMiddleware) Serve(ctx iris.Context) {
	defer func() {
		if r := recover(); r != nil {
			problem := iris.NewProblem()
			problem.Status(iris.StatusInternalServerError)
			problem.Detail(fmt.Sprintf("%+v", r))

			ctx.Problem(problem)
		}
	}()

	ctx.Next()
}
