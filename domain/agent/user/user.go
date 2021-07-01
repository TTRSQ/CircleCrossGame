package user

import (
	"fmt"

	"github.com/TTRSQ/CircleCrossGame/domain/agent"
	"github.com/TTRSQ/CircleCrossGame/domain/constants"
	"github.com/TTRSQ/CircleCrossGame/domain/game/action"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
)

// ユーザー(console)
type user struct {
	symbol constants.Symbol
}

func Get(symbol constants.Symbol) agent.Agent {
	return &user{
		symbol: symbol,
	}
}

func (u *user) NextAction(board board.Board) (*action.Item, error) {
	fmt.Println("input:x y")
	var x, y int
	fmt.Scanf("&d &d", &x, &y)
	return action.NewItem(x, y, u.symbol)
}

func (u *user) Symbol() constants.Symbol {
	return u.symbol
}
