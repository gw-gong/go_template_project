package component02

import "fmt"

type Component02er interface {
	Function01()
	Function02()
}

type component02 struct {
	field01 string
	field02 string
}

func NewComponent02(field01 string, field02 string) Component02er {
	return &component02{
		field01: field01,
		field02: field02,
	}
}

func (c *component02) Function01() {
	fmt.Println("Function01", c.field01)
}

func (c *component02) Function02() {
	fmt.Println("Function02", c.field02)
}
