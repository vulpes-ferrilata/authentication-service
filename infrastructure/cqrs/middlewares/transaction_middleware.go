package middlewares

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/cqrs"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTransactionMiddleware(db *mongo.Database) *TransactionMiddleware {
	return &TransactionMiddleware{
		db: db,
	}
}

type TransactionMiddleware struct {
	db *mongo.Database
}

func (t TransactionMiddleware) CommandHandlerMiddleware() cqrs.CommandMiddlewareFunc {
	return func(commandHandlerFunc cqrs.CommandHandlerFunc[any]) cqrs.CommandHandlerFunc[any] {
		return func(ctx context.Context, command any) error {
			session, err := t.db.Client().StartSession()
			if err != nil {
				return errors.WithStack(err)
			}

			if _, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
				if err := commandHandlerFunc(sessCtx, command); err != nil {
					return nil, errors.WithStack(err)
				}

				return nil, nil
			}); err != nil {
				return errors.WithStack(err)
			}

			return nil
		}
	}
}
