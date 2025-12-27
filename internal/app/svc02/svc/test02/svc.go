package test02

import (
	"context"

	"github.com/gw-gong/go-template-project/api/rpc/svc02"

	"github.com/gw-gong/gwkit-go/log"
)

type Test02Svc struct {
	svc02.UnimplementedTest02ServiceServer
}

func NewTest02Svc() *Test02Svc {
	return &Test02Svc{}
}

func (s *Test02Svc) TestFunc(ctx context.Context, req *svc02.Test02Request) (*svc02.Test02Response, error) {
	log.Infoc(ctx, "TestFunc", log.Str("field01", req.Field01), log.Str("field02", req.Field02))
	return &svc02.Test02Response{
		Field01: "test01",
		Field02: "test02",
	}, nil
}
