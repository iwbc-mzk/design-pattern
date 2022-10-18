package cellphone

import (
	"fmt"
)

type CellPhone struct {
	owner string
}

func newCellPhone(owner string) CellPhone {
	fmt.Printf("%sのスマホを作成しました。\n", owner)
	a := CellPhone{owner: owner}
	return a
}

func (i CellPhone) Use() {
	fmt.Printf("%sのスマホを使います。\n", i.owner)
}

func (i *CellPhone) GetOwner() string {
	return i.owner
}
