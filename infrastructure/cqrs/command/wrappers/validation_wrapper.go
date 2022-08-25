package wrappers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command"
)

func NewValidationWrapper[Command any](validate *validator.Validate, nextHandler command.CommandHandler[Command]) command.CommandHandler[Command] {
	return &validationWrapper[Command]{
		validate:    validate,
		nextHandler: nextHandler,
	}
}

type validationWrapper[Command any] struct {
	validate    *validator.Validate
	nextHandler command.CommandHandler[Command]
}

func (v validationWrapper[Command]) Handle(ctx context.Context, command Command) error {
	if err := v.validate.StructCtx(ctx, command); err != nil {
		return errors.WithStack(err)
	}

	if err := v.nextHandler.Handle(ctx, command); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
