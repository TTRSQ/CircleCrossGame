package display

import "github.com/TTRSQ/CircleCrossGame/domain/game/board"

type Display interface {
	Disp(board board.Board) error
}
