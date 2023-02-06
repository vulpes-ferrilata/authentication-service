package middlewares

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/cqrs"
)

func NewValidationMiddleware(validate *validator.Validate) *ValidationMiddleware {
	return &ValidationMiddleware{
		validate: validate,
	}
}

type ValidationMiddleware struct {
	validate *validator.Validate
}

func (v ValidationMiddleware) CommandHandlerMiddleware() cqrs.CommandMiddlewareFunc {
	return func(commandHandlerFunc cqrs.CommandHandlerFunc[any]) cqrs.CommandHandlerFunc[any] {
		return func(ctx context.Context, command any) error {
			if err := v.validate.StructCtx(ctx, command); err != nil {
				return errors.WithStack(err)
			}

			if err := commandHandlerFunc(ctx, command); err != nil {
				return errors.WithStack(err)
			}

			return nil
		}
	}
}

func (v ValidationMiddleware) QueryHandlerMiddleware() cqrs.QueryMiddlewareFunc {
	return func(commandHandlerFunc cqrs.QueryHandlerFunc[any, any]) cqrs.QueryHandlerFunc[any, any] {
		return func(ctx context.Context, command any) (interface{}, error) {
			if err := v.validate.StructCtx(ctx, command); err != nil {
				return nil, errors.WithStack(err)
			}

			result, err := commandHandlerFunc(ctx, command)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			return result, nil
		}
	}
}
