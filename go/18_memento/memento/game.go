package memento

import (
	"fmt"
)

// 時間や操作によって変化するような状態を保持するオブジェクト
// その状態をMementoに保存するメソッドと、状態を復元するメソッドを実装する
type Game struct {
	isFirstPlayer bool
	state [3]([3]string)
}

func NewGame() *Game {
	row := [3]string{"-", "-", "-"}
	state := [3]([3]string){row, row, row}
	return &Game{isFirstPlayer: true, state: state}
}

func (g *Game) setState(state [3]([3]string)) {
	g.state = state
}

func (g *Game) getState() [3]([3]string) {
	return g.state
}

// 現在の状態をMementoに保存する。
func (g *Game) backup() *memento {
	var game originator = g
	var s memento = newSnapshot(&game)
	return &s
}

func (g *Game) SetMark(row, column int) error {
	if row > 3 || column > 3 {
		return fmt.Errorf("invalid index")
	}
	if g.isFirstPlayer {
		g.state[row - 1][column - 1] = "o"
	} else {
		g.state[row - 1][column - 1] = "x"
	}
	g.isFirstPlayer = !g.isFirstPlayer
	return nil
}

func (g *Game) Show() {
	for _, v := range g.state {
		fmt.Println(v)
	}
}