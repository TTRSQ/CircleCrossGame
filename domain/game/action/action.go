package action

import (
	"errors"

	"github.com/TTRSQ/CircleCrossGame/domain/game/constants"
)

type Item struct {
	x      int
	y      int
	simbol constants.NodeStatus
}

func NewItem(x, y int, simbol constants.NodeStatus) (*Item, error) {
	if x < 0 || 2 < x {
		return nil, errors.New("x: out of range")
	}
	if y < 0 || 2 < y {
		return nil, errors.New("y: out of range")
	}
	if simbol == constants.NONE {
		return nil, errors.New("invalid simbol")
	}
	return &Item{
		x:      x,
		y:      y,
		simbol: simbol,
	}, nil
}

func (i *Item) X() int {
	return i.x
}

func (i *Item) Y() int {
	return i.y
}

func (i *Item) Simbol() constants.NodeStatus {
	return i.simbol
}
