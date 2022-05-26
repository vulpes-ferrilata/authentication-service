package context

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type transactionContextKey struct{}

func WithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, transactionContextKey{}, tx)
}

func GetTransaction(ctx context.Context) (*gorm.DB, error) {
	tx, ok := ctx.Value(transactionContextKey{}).(*gorm.DB)
	if !ok {
		return nil, errors.Wrap(ErrContextValueNotFound, "transaction")
	}
	return tx, nil
}
