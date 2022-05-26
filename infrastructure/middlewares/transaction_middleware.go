package middlewares

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/infrastructure/bus"
	infrastructure_context "github.com/VulpesFerrilata/authentication-service/infrastructure/context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewTransactionMiddleware(db *gorm.DB) *TransactionMiddleware {
	return &TransactionMiddleware{
		db: db,
	}
}

type TransactionMiddleware struct {
	db *gorm.DB
}

func (t TransactionMiddleware) WrapCommandHandler(next bus.CommandHandlerFunc) bus.CommandHandlerFunc {
	return func(ctx context.Context, command interface{}) error {
		tx := t.db.WithContext(ctx).Begin()

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r)
			}
		}()

		ctx = infrastructure_context.WithTransaction(ctx, tx)

		if err := next(ctx, command); err != nil {
			tx.Rollback()
			return errors.WithStack(err)
		}

		if err := tx.Commit().Error; err != nil {
			return errors.WithStack(err)
		}

		return nil
	}
}

func (t TransactionMiddleware) WrapQueryHandler(next bus.QueryHandlerFunc) bus.QueryHandlerFunc {
	return func(ctx context.Context, query interface{}) (interface{}, error) {
		tx := t.db.WithContext(ctx).Begin()

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r)
			}
		}()

		ctx = infrastructure_context.WithTransaction(ctx, tx)

		result, err := next(ctx, query)
		if err != nil {
			tx.Rollback()
			return nil, errors.WithStack(err)
		}

		if err := tx.Commit().Error; err != nil {
			return nil, errors.WithStack(err)
		}

		return result, nil
	}
}
