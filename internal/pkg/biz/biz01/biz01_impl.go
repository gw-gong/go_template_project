package biz01

import (
	"context"

	"github.com/gw-gong/gwkit-go/log"
)

type biz01 struct {
	field01 string
	field02 string
}

type Biz01Options struct {
	Field01 string `yaml:"field01" mapstructure:"field01"`
	Field02 string `yaml:"field02" mapstructure:"field02"`
}

func NewBiz01(options *Biz01Options) Biz01 {
	return &biz01{
		field01: options.Field01,
		field02: options.Field02,
	}
}

func (c *biz01) Function01(ctx context.Context) {
	log.Infoc(ctx, "Function01", log.Str("field01", c.field01))
}

func (c *biz01) Function02(ctx context.Context) {
	log.Infoc(ctx, "Function02", log.Str("field02", c.field02))
}
