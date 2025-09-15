package component02

import (
	"context"

	"github.com/gw-gong/gwkit-go/log"
)

type component02er struct {
	field01 string
	field02 string
}

type Component02erOptions struct {
	Field01 string
	Field02 string
}

func NewComponent02er(options Component02erOptions) (Component02er, func()) {
	return &component02er{
			field01: options.Field01,
			field02: options.Field02,
		}, func() {
			log.Info("Component02er cleanup")
		}
}

func (c *component02er) Function01(ctx context.Context) {
	log.Infoc(ctx, "Function01", log.Str("field01", c.field01))
}

func (c *component02er) Function02(ctx context.Context) {
	log.Infoc(ctx, "Function02", log.Str("field02", c.field02))
}
