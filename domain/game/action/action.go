package action

import (
	"errors"

	"github.com/TTRSQ/CircleCrossGame/domain/constants"
)

type Item struct {
	x      int
	y      int
	symbol constants.Symbol
}

func NewItem(x, y int, symbol constants.Symbol) (*Item, error) {
	if x < 0 || 2 < x {
		return nil, errors.New("x: out of range")
	}
	if y < 0 || 2 < y {
		return nil, errors.New("y: out of range")
	}
	return &Item{
		x:      x,
		y:      y,
		symbol: symbol,
	}, nil
}

func (i *Item) X() int {
	return i.x
}

func (i *Item) Y() int {
	return i.y
}

func (i *Item) Symbol() constants.Symbol {
	return i.symbol
}
