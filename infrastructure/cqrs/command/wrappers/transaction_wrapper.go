package wrappers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTransactionWrapper[Command any](db *mongo.Database, nextHandler command.CommandHandler[Command]) command.CommandHandler[Command] {
	return &transactionWrapper[Command]{
		db:          db,
		nextHandler: nextHandler,
	}
}

type transactionWrapper[Command any] struct {
	db          *mongo.Database
	nextHandler command.CommandHandler[Command]
}

func (t transactionWrapper[Command]) Handle(ctx context.Context, command Command) error {
	session, err := t.db.Client().StartSession()
	if err != nil {
		return errors.WithStack(err)
	}

	if _, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		if err := t.nextHandler.Handle(sessCtx, command); err != nil {
			return nil, errors.WithStack(err)
		}

		return nil, nil
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
