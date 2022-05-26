package bus

import (
	"context"
)

type CommandHandler interface {
	GetCommand() interface{}
	Handle(ctx context.Context, command interface{}) error
}
