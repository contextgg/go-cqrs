package cqrs

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

type FakeCommand struct {
}
type Fake2Command struct {
}

type FakeCommandHandler struct {
	mock.Mock
}

func (h *FakeCommandHandler) Handle(ctx context.Context, cmd Command) error {
	args := h.Called(ctx, cmd)
	return args.Error(0)
}

func TestBus(t *testing.T) {
	ctx := context.TODO()
	cmd1 := &FakeCommand{}

	handler := new(FakeCommandHandler)
	handler.On("Handle", mock.Anything, cmd1).Return(nil)

	bus := NewBus(
		BusCommandHandler(cmd1, handler),
	)

	if err := bus.Send(ctx, cmd1); err != nil {
		t.Error(err)
		return
	}

	cmd2 := &Fake2Command{}
	if err := bus.Send(ctx, cmd2); !errors.Is(err, ErrCommandHandlerNotFound) {
		t.Error("Shouldn't be found", err)
		return
	}

	handler.AssertExpectations(t)
}

func TestAggregate(t *testing.T) {
}
