package main

// Invoker(起動者)
// Commandインタフェースで宣言されたexecuteを呼び出し、命令を実行する
// Invokerは具象Commandには依存しない。
type App struct {
	editor *Editor
	histroy []Command
}

func NewApp(editor *Editor) *App {
	return &App{editor: editor}
}

func (a *App) ExecuteCommand(command Command) {
	a.histroy = append(a.histroy, command)
	command.execute()
}

func (a *App) Input(text string) {
	command := NewInputCommand(text, a.editor)
	a.executeCommand(command)
}

func (a *App) Undo() {
	if len(a.histroy) > 0 {
		command := a.histroy[len(a.histroy) - 1]
		command.execute()
		a.histroy = a.histroy[:len(a.histroy) - 1]
	}
}

func (a *App) Clear() {
	command := NewClearCommand(a.editor)
	a.executeCommand(command)
}

func (a *App) executeCommand(command Command) {
	command.execute()
	a.histroy = append(a.histroy, command)
}