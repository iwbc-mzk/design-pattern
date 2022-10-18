package idcard

import (
	"fmt"
)

type IDCard struct {
	owner string
}

// idcardパッケージ以外から生成できないように定義する
// →生成は必ずIDCardFactory経由で行うようにする
func newIDCard(owner string) IDCard {
	fmt.Println(owner + "のカードを作ります。")
	i := IDCard{owner: owner}
	return i
}

func (i IDCard) Use() {
	fmt.Println(i.owner + "のカードを使います。")
}

func (i *IDCard) GetOwner() string {
	return i.owner
}
