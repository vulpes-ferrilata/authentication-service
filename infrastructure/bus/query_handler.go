package bus

import (
	"context"
)

type QueryHandler interface {
	GetQuery() interface{}
	Handle(ctx context.Context, query interface{}) (interface{}, error)
}
