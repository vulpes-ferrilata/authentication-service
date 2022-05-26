package bus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
)

type CommandBus interface {
	Register(commandHandlers ...CommandHandler) error
	Use(commandHandlerWrappers ...CommandHandlerWrapper)
	Execute(ctx context.Context, command interface{}) error
}

func NewCommandBus() CommandBus {
	commandBus := &commandBus{
		commandHandlerFuncs:    make(map[string]CommandHandlerFunc),
		commandHandlerWrappers: make([]CommandHandlerWrapper, 0),
	}

	return commandBus
}

type commandBus struct {
	commandHandlerFuncs    map[string]CommandHandlerFunc
	commandHandlerWrappers []CommandHandlerWrapper
}

func (c *commandBus) addCommandHandler(commandHandler CommandHandler) error {
	command := commandHandler.GetCommand()
	commandName := reflect.TypeOf(command).String()

	if _, ok := c.commandHandlerFuncs[commandName]; ok {
		return errors.Errorf("command (%s) is already assigned", commandName)
	}

	c.commandHandlerFuncs[commandName] = commandHandler.Handle

	return nil
}

func (c *commandBus) Register(commandHandlers ...CommandHandler) error {
	for _, commandHandler := range commandHandlers {
		if err := c.addCommandHandler(commandHandler); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (c *commandBus) Use(commandHandlerWrappers ...CommandHandlerWrapper) {
	c.commandHandlerWrappers = append(c.commandHandlerWrappers, commandHandlerWrappers...)
}

func (c commandBus) Execute(ctx context.Context, command interface{}) error {
	commandName := reflect.TypeOf(command).String()

	commandHandlerFunc, ok := c.commandHandlerFuncs[commandName]
	if !ok {
		return errors.Errorf("handler not found for command (%s)", commandName)
	}

	for i := len(c.commandHandlerWrappers) - 1; i >= 0; i-- {
		commandHandlerFunc = c.commandHandlerWrappers[i](commandHandlerFunc)
	}

	if err := commandHandlerFunc(ctx, command); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
