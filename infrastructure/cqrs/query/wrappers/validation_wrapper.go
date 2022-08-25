package wrappers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/query"
)

func NewValidationWrapper[Query any, Result any](validate *validator.Validate, nextHandler query.QueryHandler[Query, Result]) query.QueryHandler[Query, Result] {
	return &validationWrapper[Query, Result]{
		validate:    validate,
		nextHandler: nextHandler,
	}
}

type validationWrapper[Query any, Result any] struct {
	validate    *validator.Validate
	nextHandler query.QueryHandler[Query, Result]
}

func (v validationWrapper[Query, Result]) Handle(ctx context.Context, query Query) (Result, error) {
	var result Result

	if err := v.validate.StructCtx(ctx, query); err != nil {
		return result, errors.WithStack(err)
	}

	result, err := v.nextHandler.Handle(ctx, query)
	if err != nil {
		return result, errors.WithStack(err)
	}

	return result, nil
}
