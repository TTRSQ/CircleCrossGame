package board

import (
	"errors"

	"github.com/TTRSQ/CircleCrossGame/domain/game/action"
	"github.com/TTRSQ/CircleCrossGame/domain/game/constants"
)

/*
task:
- 状態管理
- 遷移可能状態列挙
- 現状調査
*/

const (
	WIDTH  = 3
	HEIGHT = 3
)

type Board []constants.NodeStatus

func NewBoard() Board {
	return Board{
		constants.NONE, constants.NONE, constants.NONE,
		constants.NONE, constants.NONE, constants.NONE,
		constants.NONE, constants.NONE, constants.NONE,
	}
}

func (b Board) Act(act action.Item) error {
	if b[act.Y()*WIDTH+act.X()] != constants.NONE {
		return errors.New("already placed")
	}
	b[act.Y()*WIDTH+act.X()] = act.Simbol()
	return nil
}

func lineBinaryList() []int {
	return []int{
		448, // 111,000,000
		56,  // 000,111,000
		7,   // 000,000,111
		292, // 100,100,100
		146, // 010,010,010
		73,  // 001,001,001
		273, // 100,010,001
		84,  // 001,010,100
	}
}

func (b Board) Binary(simbol constants.NodeStatus) (int, error) {
	if simbol == constants.NONE {
		return 0, errors.New("invald simbol")
	}
	stamp := []int{
		512, 256, 128,
		64, 32, 16,
		8, 4, 2,
	}
	sum := 0
	for i := range b {
		if b[i] == simbol {
			sum += stamp[i]
		}
	}
	return sum, nil
}

func (b Board) CountLine(simbol constants.NodeStatus) (int, error) {
	if simbol == constants.NONE {
		return 0, errors.New("invald simbol")
	}

	binary, _ := b.Binary(simbol)
	count := 0
	for _, line := range lineBinaryList() {
		if binary&line == line {
			count += 1
		}
	}
	return count, nil
}
