package memento

type snapshot struct {
	originator *originator
	state [3]([3]string)
}

func newSnapshot(o *originator) *snapshot {
	return &snapshot{
		originator: o,
		state: (*o).getState(),
	}
}

// ある時点の状態でOriginatorの状態を復元する。
func (s *snapshot) restore() {
	(*(s.originator)).setState(s.state)
}