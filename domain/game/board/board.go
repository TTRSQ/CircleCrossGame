package board

import "github.com/TTRSQ/CircleCrossGame/domain/game/action"

/*
task:
- 状態管理
- 遷移可能状態列挙
- 現状調査
*/

type Board struct {
	status [][]int
}

func (b *Board) Act(act action.Item) error {
	return nil
}
