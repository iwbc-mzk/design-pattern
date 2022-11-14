package main

var open *Open

type Open struct {}

// Stateオブジェクトはインスタンスフィールドを持たない。
// そのためシングルトンパターンを利用して実装することが可能。(必須ではない)
func init() {
	open = &Open{}
}

// 今回の実装例ではApprove()...の引数にContextを渡しているが、Stateオブジェクト生成時にContextの参照を保持するようにもできる。
// その場合はシングルトンは使えない。
func GetOpneInstance() *Open {
	return open
}

// Contesxtオブジェクトの状態変更メソッドを利用して、現在の状態に応じた適切な状態変更を行う。
// 状態遷移の責任を具象Stateオブジェクトが担うことになる。
// そのため他の具象Stateオブジェクトへの依存が必要となる。
func (o *Open) Approve(ticket *Ticket) {
	// Open -> In Progress
	ticket.ChangeState(GetInProgressInstance())
}

func (o *Open) Reject(ticket *Ticket) {
	// Open -> Close
	ticket.ChangeState(GetCloseInstance())
}

func (o *Open) Close(ticket *Ticket) {
	// Open -> Close
	ticket.ChangeState(GetCloseInstance())
}

func (o *Open) GetState() string {
	return "Open"
}