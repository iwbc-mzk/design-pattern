package memento

type CareTaker struct {
	game *Game
	history []*memento
}

func NewCareTaker(game *Game) *CareTaker {
	return &CareTaker{game: game}
}

func (c *CareTaker) MakeBackup() {
	c.history = append(c.history, c.game.createMemento())
}

func (c *CareTaker) Undo() {
	if memento := c.popHistory(); memento != nil {
		c.game.restoreMemento(memento)
	}
}

func (c *CareTaker) popHistory() *memento {
	if len(c.history) == 0 {
		return nil
	}
	lastIndex := len(c.history) - 1
	memento := c.history[lastIndex]
	c.history = c.history[:lastIndex]
	return memento
}