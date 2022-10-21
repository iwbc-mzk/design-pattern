package main

import (
	"fmt"
	"strings"
)

// 子要素を持つことができる複雑なComponentを表現する
type Table struct {
	components []Component
	values []string
}

func NewTable(s ...string) *Table {
	return &Table{values: s, components: []Component{}}
}

func (t *Table) show(indent int) {
	fmt.Printf("%s%v\n", strings.Repeat("--", indent), t.values)
	for _, composite := range t.components {
		composite.show(indent + 1)
	}
}

func (t *Table) add(c ...Component) {
	t.components = append(t.components, c...)
}