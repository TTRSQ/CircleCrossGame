package agent

import (
	"github.com/TTRSQ/CircleCrossGame/domain/constants"
	"github.com/TTRSQ/CircleCrossGame/domain/game/action"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
)

/*
## task
- 入力方法の抽象化
*/

type Agent interface {
	NextAction(board board.Board) (*action.Item, error)
	Symbol() constants.Symbol
}
