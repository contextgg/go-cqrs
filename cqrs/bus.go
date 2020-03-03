package cqrs

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

var ErrCommandHandlerNotFound = errors.New("Command handler not found")

type Bus struct {
	handlers map[reflect.Type]CommandHandler
}

func (b Bus) Send(ctx context.Context, cmd Command) error {
	t := reflect.TypeOf(cmd)

	// get the name!
	for key, handler := range b.handlers {
		if t == key {
			return handler.Handle(ctx, cmd)
		}
	}

	return fmt.Errorf("Command %v - %w", t.Name(), ErrCommandHandlerNotFound)
}

func NewBus(cfgs ...BusConfig) *Bus {
	b := &Bus{
		handlers: make(map[reflect.Type]CommandHandler),
	}
	for _, c := range cfgs {
		c(b)
	}
	return b
}

type BusConfig func(*Bus)

func BusCommandHandler(cmd Command, handler CommandHandler) BusConfig {
	t := reflect.TypeOf(cmd)

	return func(b *Bus) {
		b.handlers[t] = handler
	}
}
