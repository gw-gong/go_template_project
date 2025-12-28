package test01

import (
	"context"

	"github.com/gw-gong/go-template-project/api/rpc/svc02"

	"github.com/gw-gong/gwkit-go/log"
)

type Test01Svc struct {
	svc02.UnimplementedTest01ServiceServer
}

func NewTest01Svc() *Test01Svc {
	return &Test01Svc{}
}

func (s *Test01Svc) TestFunc(ctx context.Context, req *svc02.Test01Request) (*svc02.Test01Response, error) {
	log.Infoc(ctx, "TestFunc", log.Str("field01", req.Field01), log.Str("field02", req.Field02))
	return &svc02.Test01Response{
		Field01: "test01",
		Field02: "test02",
	}, nil
}
