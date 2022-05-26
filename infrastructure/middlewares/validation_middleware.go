package middlewares

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/infrastructure/bus"
	app_errors "github.com/VulpesFerrilata/authentication-service/infrastructure/errors"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func NewValidationMiddleware(validate *validator.Validate) *ValidationMiddleware {
	return &ValidationMiddleware{
		validate: validate,
	}
}

type ValidationMiddleware struct {
	validate *validator.Validate
}

func (v ValidationMiddleware) WrapCommandHandler(next bus.CommandHandlerFunc) bus.CommandHandlerFunc {
	return func(ctx context.Context, command interface{}) error {
		err := v.validate.StructCtx(ctx, command)
		if validationErrs, ok := errors.Cause(err).(validator.ValidationErrors); ok {
			return app_errors.NewValidationError(validationErrs...)
		}
		if err != nil {
			return errors.WithStack(err)
		}

		if err := next(ctx, command); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}
}

func (v ValidationMiddleware) WrapQueryHandler(next bus.QueryHandlerFunc) bus.QueryHandlerFunc {
	return func(ctx context.Context, query interface{}) (interface{}, error) {
		err := v.validate.StructCtx(ctx, query)
		if validationErrs, ok := errors.Cause(err).(validator.ValidationErrors); ok {
			return nil, app_errors.NewValidationError(validationErrs...)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}

		result, err := next(ctx, query)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return result, nil
	}
}
