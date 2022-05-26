package middlewares

import (
	"context"

	infrastructure_context "github.com/VulpesFerrilata/authentication-service/infrastructure/context"
	app_errors "github.com/VulpesFerrilata/authentication-service/infrastructure/errors"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func NewErrorHandlerMiddleware() *ErrorHandlerMiddleware {
	return &ErrorHandlerMiddleware{}
}

type ErrorHandlerMiddleware struct{}

func (e ErrorHandlerMiddleware) WrapCall(f client.CallFunc) client.CallFunc {
	return func(ctx context.Context, node *registry.Node, request client.Request, response interface{}, opts client.CallOptions) error {
		err := f(ctx, node, request, response, opts)
		if stt, ok := status.FromError(errors.Cause(err)); ok && stt != nil {
			return app_errors.NewStatusError(stt)
		}
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	}
}

func (e ErrorHandlerMiddleware) WrapHandler(f server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := f(ctx, req, rsp)
		if detailErr, ok := errors.Cause(err).(app_errors.DetailError); ok {
			err = app_errors.NewBusinessRuleError(detailErr)
		}
		if grpcErr, ok := errors.Cause(err).(app_errors.GrpcError); ok {
			translator, err := infrastructure_context.GetTranslator(ctx)
			if err != nil {
				return errors.WithStack(err)
			}

			stt, err := grpcErr.Status(translator, req.Service())
			if err != nil {
				return errors.WithStack(err)
			}

			return stt.Err()
		}
		if err != nil {
			errors.WithStack(err)
		}

		return nil
	}
}

func (e ErrorHandlerMiddleware) Handle(ctx iris.Context, err error) {
	if detailErr, ok := errors.Cause(err).(app_errors.DetailError); ok {
		err = app_errors.NewBusinessRuleError(detailErr)
	}

	if webErr, ok := errors.Cause(err).(app_errors.WebError); ok {
		translator, err := infrastructure_context.GetTranslator(ctx.Request().Context())
		if err != nil {
			e.Handle(ctx, err)
			return
		}

		problem, err := webErr.Problem(translator)
		if err != nil {
			e.Handle(ctx, err)
			return
		}
		ctx.Problem(problem)
		return
	}

	problem := iris.NewProblem()
	problem.Status(iris.StatusInternalServerError)
	problem.Title("something went wrong")
	problem.Detail(err.Error())
	ctx.Problem(problem)
}
