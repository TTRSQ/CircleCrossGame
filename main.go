package main

import (
	"fmt"
	"log"

	"github.com/TTRSQ/CircleCrossGame/domain/agent/randcp"
	"github.com/TTRSQ/CircleCrossGame/domain/agent/user"
	"github.com/TTRSQ/CircleCrossGame/domain/constants"
	"github.com/TTRSQ/CircleCrossGame/domain/game"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
	"github.com/TTRSQ/CircleCrossGame/domain/game/display/console"
)

func main() {
	manager, err := game.NewManager(
		user.Get(constants.CIRCLE, "you"),
		randcp.Get(constants.CROSS, "cp"),
		board.NewBoard(),
		console.Get(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = manager.Play()

	if err == nil {
		winner := manager.Winner()
		if winner == nil {
			fmt.Println("drow")
		} else {
			fmt.Println("winner: " + winner.Name())
		}
	} else {
		log.Fatalln(err)
	}
}
