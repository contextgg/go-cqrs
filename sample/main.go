package main

import (
	"context"

	"github.com/contextgg/go-cqrs/cqrs"
)

func main() {
	ctx := context.Background()
	var handlers LeaderboardCommandHandlers
	var cmd CreateLeaderboard

	bus := cqrs.NewBus(
		cqrs.BusCommandHandler(cmd, cqrs.CommandHandlerFunc(handlers.CreateLeaderboard)),
	)

	if err := bus.Send(ctx, cmd); err != nil {
		panic(err)
	}

	print("CMD ", cmd.Name)
}
