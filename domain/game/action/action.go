package action

import (
	"errors"

	"github.com/TTRSQ/CircleCrossGame/domain/game/constants"
)

type Item struct {
	x      int
	y      int
	status constants.NodeStatus
}

func NewItem(x, y int, staus constants.NodeStatus) (*Item, error) {
	if x < 0 || 2 < x {
		return nil, errors.New("x: out of range")
	}
	if y < 0 || 2 < y {
		return nil, errors.New("y: out of range")
	}
}
