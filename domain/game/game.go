package geme

import (
	"github.com/TTRSQ/CircleCrossGame/domain/agent"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
)

/*
game のマネジメント
*/

type Manager struct {
	FirstAgent  agent.Agent
	SecondAgent agent.Agent
	board       board.Board
}
