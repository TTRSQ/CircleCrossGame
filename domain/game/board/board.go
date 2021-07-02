package board

import (
	"errors"

	"github.com/TTRSQ/CircleCrossGame/domain/constants"
	"github.com/TTRSQ/CircleCrossGame/domain/game/action"
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

type Board struct {
	status []*constants.Symbol
}

func NewBoard() Board {
	return Board{
		status: []*constants.Symbol{
			nil, nil, nil,
			nil, nil, nil,
			nil, nil, nil,
		},
	}
}

func (b *Board) Act(act action.Item) error {
	if b.status[act.Y()*WIDTH+act.X()] != nil {
		return errors.New("already placed")
	}
	symbol := act.Symbol()
	b.status[act.Y()*WIDTH+act.X()] = &symbol
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

func (b *Board) binary(symbol constants.Symbol) (int, error) {
	stamp := []int{
		256, 128, 64,
		32, 16, 8,
		4, 2, 1,
	}
	sum := 0
	for i := range b.status {
		if b.status[i] != nil && *b.status[i] == symbol {
			sum += stamp[i]
		}
	}
	return sum, nil
}

func (b *Board) CountLine(symbol constants.Symbol) (int, error) {
	binary, _ := b.binary(symbol)
	count := 0
	for _, line := range lineBinaryList() {
		if binary&line == line {
			count += 1
		}
	}
	return count, nil
}

func (b *Board) CanPutPoints() [][]int {
	ret := [][]int{}
	for i := range b.status {
		if b.status[i] == nil {
			ret = append(ret, []int{i % WIDTH, i / WIDTH})
		}
	}
	return ret
}

func (b *Board) Status() [][]*constants.Symbol {
	ret := [][]*constants.Symbol{}

	for y := 0; y < HEIGHT; y++ {
		row := []*constants.Symbol{}
		for x := 0; x < WIDTH; x++ {
			row = append(row, b.status[y*WIDTH+x])
		}
		ret = append(ret, row)
	}

	return ret
}

func (b *Board) CountSymbol(symbol constants.Symbol) int {
	sum := 0
	for i := range b.status {
		if b.status[i] != nil && *b.status[i] == symbol {
			sum++
		}
	}
	return sum
}
