package memento

type CareTaker struct {
	originator originator
	history []*memento
}

func NewCareTaker(originator originator) *CareTaker {
	return &CareTaker{originator: originator}
}

func (c *CareTaker) MakeBackup() {
	c.history = append(c.history, c.originator.backup())
}

func (c *CareTaker) Undo() {
	if backup := c.popHistory(); backup != nil {
		(*backup).restore()
	}
}

func (c *CareTaker) popHistory() *memento {
	if len(c.history) == 0 {
		return nil
	}
	lastIndex := len(c.history) - 1
	backup := c.history[lastIndex]
	c.history = c.history[:lastIndex]
	return backup
}
