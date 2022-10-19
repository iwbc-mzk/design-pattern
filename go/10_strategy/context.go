package main

type Context struct {
	exchanger Exchanger
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) Exchange(yen float64) float64 {
	return c.exchanger.Exchange(yen)
}

func (c *Context) Rate() float64 {
	return c.exchanger.GetRate()
}

func (c *Context) SetExchanger(e Exchanger) {
	c.exchanger = e
}