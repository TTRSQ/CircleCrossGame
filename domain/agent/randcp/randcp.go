package randcp

import (
	"fmt"
	"math/rand"

	"github.com/TTRSQ/CircleCrossGame/domain/agent"
	"github.com/TTRSQ/CircleCrossGame/domain/constants"
	"github.com/TTRSQ/CircleCrossGame/domain/game/action"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
)

// ランダムな打ち手(cp)
type randcp struct {
	symbol constants.Symbol
	name   string
}

func Get(symbol constants.Symbol, name string) agent.Agent {
	return &randcp{
		symbol: symbol,
		name:   name,
	}
}

func (r *randcp) NextAction(board board.Board) (*action.Item, error) {
	canPutPoints := board.CanPutPoints()
	selection := rand.Int() % len(canPutPoints)
	fmt.Println(r.name + " placed " + fmt.Sprintf(
		"(%d %d)",
		canPutPoints[selection][0],
		canPutPoints[selection][1],
	))
	return action.NewItem(canPutPoints[selection][0], canPutPoints[selection][1], r.symbol)
}

func (r *randcp) Symbol() constants.Symbol {
	return r.symbol
}

func (r *randcp) Name() string {
	return r.name
}
