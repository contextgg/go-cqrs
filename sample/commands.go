package main

import (
	"context"
	"errors"

	"github.com/contextgg/go-cqrs/cqrs"
)

var ErrWrongCommandType = errors.New("Wrong command type")

// LeaderboardCommands
type LeaderboardCommandHandlers struct {
	Store cqrs.AggregateStore
}

func (h *LeaderboardCommandHandlers) CreateLeaderboard(ctx context.Context, cmd cqrs.Command) error {
	cl, ok := cmd.(*CreateLeaderboard)
	if !ok {
		return ErrWrongCommandType
	}

	// get the
	a, err := h.Store.Load(cl.ID)
	if err != nil {
		return err
	}

	// run our logic here.

	var evts []cqrs.Event
	h.Store.AppendEvents(ctx, evts)

	return nil
}

// CreateLeaderboard command
type CreateLeaderboard struct {
	ID   string
	Name string
}
