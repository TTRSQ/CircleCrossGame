package randcp

import (
	"math/rand"

	"github.com/TTRSQ/CircleCrossGame/domain/agent"
	"github.com/TTRSQ/CircleCrossGame/domain/constants"
	"github.com/TTRSQ/CircleCrossGame/domain/game/action"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
)

// ランダムな打ち手(cp)
type randcp struct {
	symbol constants.Symbol
}

func Get(symbol constants.Symbol) agent.Agent {
	return &randcp{
		symbol: symbol,
	}
}

func (r *randcp) NextAction(board board.Board) (*action.Item, error) {
	canPutPoints := board.CanPutPoints()
	selection := rand.Int() / len(canPutPoints)
	return action.NewItem(canPutPoints[selection][0], canPutPoints[selection][1], u.symbol)
}

func (r *randcp) Symbol() constants.Symbol {
	return u.symbol
}
