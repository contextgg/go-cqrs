package cqrs

import "context"

type AggregateCommand interface {
}

type Command interface{}

type CommandHandler interface {
	Handle(context.Context, Command) error
}

type CommandHandlerFunc func(context.Context, Command) error

func (h CommandHandlerFunc) Handle(ctx context.Context, cmd Command) error {
	return h(ctx, cmd)
}
