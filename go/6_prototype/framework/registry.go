package framework

import (
	"fmt"
)

type Registry struct {
	showcase map[string]Product
}

func NewRegistry() *Registry {
	return &Registry{showcase: map[string]Product{}}
}

func (m *Registry) Register(name string, proto Product) {
	m.showcase[name] = proto
}

func (m *Registry) Create(protoName string) (*Product, error) {
	p, ok := m.showcase[protoName]
	if !ok {
		return nil, fmt.Errorf("not registerd")
	}
	c := p.CreateClone()
	return &c, nil
}
