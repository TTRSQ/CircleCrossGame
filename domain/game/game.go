package game

import (
	"errors"
	"fmt"

	"github.com/TTRSQ/CircleCrossGame/domain/agent"
	"github.com/TTRSQ/CircleCrossGame/domain/game/board"
	"github.com/TTRSQ/CircleCrossGame/domain/game/display"
)

/*
game のマネジメント
*/

type Manager struct {
	firstAgent  agent.Agent
	secondAgent agent.Agent
	board       board.Board
	display     display.Display
}

func NewManager(first, second agent.Agent, board board.Board, display display.Display) (*Manager, error) {
	if first == nil {
		return nil, errors.New("firstAgent required")
	}
	if second == nil {
		return nil, errors.New("secondAgent required")
	}
	if first.Symbol() == second.Symbol() {
		return nil, errors.New("first and secnd must have different symbol")
	}
	if display == nil {
		return nil, errors.New("display required")
	}
	return &Manager{
		firstAgent:  first,
		secondAgent: second,
		board:       board,
		display:     display,
	}, nil
}

func (m *Manager) acter() (agent.Agent, error) {
	firstCounts := m.board.CountSymbol(m.firstAgent.Symbol())
	secondCounts := m.board.CountSymbol(m.secondAgent.Symbol())
	if firstCounts+secondCounts == board.WIDTH*board.HEIGHT {
		return nil, nil
	}
	if firstCounts == secondCounts {
		return m.firstAgent, nil
	}
	if firstCounts == secondCounts+1 {
		return m.secondAgent, nil
	}
	return nil, errors.New("symbol count invalid status" + fmt.Sprintf(
		"(%d, %d)",
		firstCounts,
		secondCounts,
	))
}

func (m *Manager) winner() (agent.Agent, error) {
	firstCounts, err := m.board.CountLine(m.firstAgent.Symbol())
	if err != nil {
		return nil, err
	}
	secondCounts, err := m.board.CountLine(m.secondAgent.Symbol())
	if err != nil {
		return nil, err
	}
	if firstCounts == 0 && secondCounts == 0 {
		return nil, nil
	}
	if firstCounts > 0 && secondCounts == 0 {
		return m.firstAgent, nil
	}
	if secondCounts > 0 && firstCounts == 0 {
		return m.secondAgent, nil
	}
	return nil, errors.New("line count invalid status" + fmt.Sprintf(
		"(%d, %d)",
		firstCounts,
		secondCounts,
	))
}

func (m *Manager) Play() error {
	for {
		winner, err := m.winner()
		if err != nil {
			return err
		}

		if winner == nil {
			// game continue
			m.display.Disp(m.board)
			acter, err := m.acter()
			if err != nil {
				return err
			}
			if acter == nil {
				// drow
				return nil
			}
			act, err := acter.NextAction(m.board)
			if err != nil {
				return err
			}
			err = m.board.Act(*act)
			if err != nil {
				return err
			}
		} else {
			// game end
			m.display.Disp(m.board)
			return nil
		}
	}
}
