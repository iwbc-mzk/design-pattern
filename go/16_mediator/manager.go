package main

type Manager struct {
	isRequesting bool
	queue        []IRequest
}

func NewManger() *Manager {
	return &Manager{}
}

func (m *Manager) canRequest(r IRequest) bool {
	if m.isRequesting {
		m.queue = append(m.queue, r)
		return false
	}
	m.isRequesting = true
	return true
}

// Componentからの処理完了通知用メソッド。
// キューから次にリクエストを送るComponentを取り出し、リクエスト処理の許可を与える。 
func (m *Manager) notifyRequestFinished() {
	if m.isRequesting {
		m.isRequesting = false
	}
	if len(m.queue) > 0 {
		firstQueue := m.queue[0]
		m.queue = m.queue[1:]
		firstQueue.permitSending()
	}
}