package component02

import (
	"context"

	"github.com/gw-gong/gwkit-go/log"
)

type component02 struct {
	field01 string
	field02 string
}

type Component02Options struct {
	Field01 string `yaml:"field01" mapstructure:"field01"`
	Field02 string `yaml:"field02" mapstructure:"field02"`
}

func NewComponent02(options Component02Options) (Component02, func()) {
	return &component02{
			field01: options.Field01,
			field02: options.Field02,
		}, func() {
			log.Info("Component02 cleanup")
		}
}

func (c *component02) Function01(ctx context.Context) {
	log.Infoc(ctx, "Function01", log.Str("field01", c.field01))
}

func (c *component02) Function02(ctx context.Context) {
	log.Infoc(ctx, "Function02", log.Str("field02", c.field02))
}
