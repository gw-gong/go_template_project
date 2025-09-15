package component01

import (
	"context"

	"github.com/gw-gong/gwkit-go/log"
)

type component01er struct {
	field01 string
	field02 string
}

type Component01erOptions struct {
	Field01 string
	Field02 string
}

func NewComponent01er(options Component01erOptions) Component01er {
	return &component01er{
		field01: options.Field01,
		field02: options.Field02,
	}
}

func (c *component01er) Function01(ctx context.Context) {
	log.Infoc(ctx, "Function01", log.Str("field01", c.field01))
}

func (c *component01er) Function02(ctx context.Context) {
	log.Infoc(ctx, "Function02", log.Str("field02", c.field02))
}
