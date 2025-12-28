//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/gw-gong/go-template-project/internal/config/svc02/localcfg"
	"github.com/gw-gong/go-template-project/internal/app/svc02/svc/test01"
	"github.com/gw-gong/go-template-project/internal/app/svc02/svc/test02"

	"github.com/gw-gong/gwkit-go/grpc/consul"
	"github.com/google/wire"
)

func InitRpcServer(config *localcfg.Config) (*RpcServer, func(), error) {
	wire.Build(
		wire.FieldsOf(
			new(*localcfg.Config),
			"ConsulAgentAddr",
		),
		consul.NewConsulClient,
		test01.NewTest01Svc,
		test02.NewTest02Svc,
		wire.Struct(new(RpcServer), "*"),
	)
	return nil, nil, nil
}
