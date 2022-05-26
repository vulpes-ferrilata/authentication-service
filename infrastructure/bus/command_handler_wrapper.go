package bus

import "context"

type CommandHandlerFunc func(context.Context, interface{}) error

type CommandHandlerWrapper func(next CommandHandlerFunc) CommandHandlerFunc
