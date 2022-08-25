package command

import (
	"context"
)

type CommandHandler[Command any] interface {
	Handle(ctx context.Context, command Command) error
}
