package component01

import "fmt"

type Component01er interface {
	Function01()
	Function02()
}

type component01 struct {
	field01 string
	field02 string
}

func NewComponent01(field01 string, field02 string) Component01er {
	return &component01{
		field01: field01,
		field02: field02,
	}
}

func (c *component01) Function01() {
	fmt.Println("Function01", c.field01)
}

func (c *component01) Function02() {
	fmt.Println("Function02", c.field02)
}
