package memento

type snapshot struct {
	game Game
	state [3]([3]string)
}

// func (s *snapshot) getState() [3]([3]string) {
// 	return s.state
// }

func (s *snapshot) restore() {
	s.game.restoreMemento()
}