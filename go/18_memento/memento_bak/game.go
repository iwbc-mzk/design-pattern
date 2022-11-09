package memento

import (
	"fmt"
)

type Game struct {
	isFirstPlayer bool
	state [3]([3]string)
}

func NewGame() *Game {
	row := [3]string{"-", "-", "-"}
	state := [3]([3]string){row, row, row}
	return &Game{isFirstPlayer: true, state: state}
}

func (g *Game) createMemento() *memento {
	return &memento{state: g.state}
}

// func (g *Game) restoreMemento(m *memento) {
// 	g.state = m.getState()
// }

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