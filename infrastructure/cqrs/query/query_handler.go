package query

import (
	"context"
)

type QueryHandler[Query any, Result any] interface {
	Handle(ctx context.Context, query Query) (Result, error)
}
