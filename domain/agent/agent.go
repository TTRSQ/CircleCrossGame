package agent

import "github.com/TTRSQ/CircleCrossGame/domain/game/action"

/*
## task
- 入力方法の抽象化
*/

type Agent interface {
	NextAction() []action.Item
}
