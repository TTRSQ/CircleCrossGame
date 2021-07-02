package console

import (
	"fmt"
	"strings"

	"github.com/TTRSQ/CircleCrossGame/domain/constants"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
	"github.com/TTRSQ/CircleCrossGame/domain/game/display"
)

type console struct{}

func Get() display.Display {
	return &console{}
}

func (c *console) Disp(board board.Board) error {
	status := board.Status()
	for _, row := range status {
		items := []string{}
		for _, elm := range row {
			items = append(items, strMapping(elm))
		}
		fmt.Println(strings.Join(items, "|"))
	}
	return nil
}

func strMapping(elm *constants.Symbol) string {
	if elm != nil {
		if *elm == constants.CIRCLE {
			return "o"

		} else if *elm == constants.CROSS {
			return "x"
		}
	}
	return " "
}
