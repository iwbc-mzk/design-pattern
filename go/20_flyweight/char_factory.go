package main

import "sync"

var instance CharFactory
var mu sync.Mutex

func init() {
	instance = CharFactory{pool: map[rune]*Char{}}
}

// Flyweightを生成・プールするファクトリ
type CharFactory struct {
	pool map[rune]*Char
}

func GetCharFactory() *CharFactory {
	return &instance
}

// Flyweight生成時に内因的状態を指定する
// プールを確認し、存在しなければ新規生成、既存ならばそのオブジェクトを返す
func (f *CharFactory) GetChar(char rune) *Char {
	// ほぼ同時にこのメソッドが呼ばれる場合を考えて排他制御をおこなう
	defer mu.Unlock()
	mu.Lock()
	if v, ok := f.pool[char]; ok {
		return v
	}

	c := NewChar(char)
	f.pool[char] = c
	return c
}