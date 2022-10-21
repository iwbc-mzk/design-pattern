package main

import (
	"fmt"
	"strings"
)

// ツリーの末端オブジェクトを表現する。
// 基本的に実際の作業はここで定義される。
type Row struct {
	values []string
}

func NewRow(s ...string) *Row {
	return &Row{values: s}
}

func (r *Row) show(indent int) {
	fmt.Printf("%s%v\n", strings.Repeat("--", indent), r.values)
}

// Componentで宣言されている場合は実装する必要あり。
// 何も行わないか、エラーを返すようにする。
// func (r *Row) add(c Component) {}