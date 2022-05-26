package bus

import "context"

type QueryHandlerFunc func(context.Context, interface{}) (interface{}, error)

type QueryHandlerWrapper func(next QueryHandlerFunc) QueryHandlerFunc
