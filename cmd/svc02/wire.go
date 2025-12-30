//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/gw-gong/boilerplate-go/internal/app/svc02/svc/test01"
	"github.com/gw-gong/boilerplate-go/internal/app/svc02/svc/test02"
	"github.com/gw-gong/boilerplate-go/internal/config/svc02/localcfg"

	"github.com/google/wire"
	"github.com/gw-gong/gwkit-go/grpc/consul"
	"github.com/gw-gong/gwkit-go/hotcfg"
)

var ConfigSet = wire.NewSet(
	localcfg.NewConfig,
	wire.FieldsOf(
		new(*localcfg.Config),
		"ConsulAgentAddr",
	),
	hotcfg.NewHotLoaderManager,
)

var InfraSet = wire.NewSet(
	consul.NewConsulClient,
)

var BizSet = wire.NewSet(
	test01.NewTest01Svc,
	test02.NewTest02Svc,
)

var RpcServerSet = wire.NewSet(
	wire.Struct(new(RpcServer), "*"),
)

func InitRpcServer(cfgOption *hotcfg.LocalConfigOption) (*RpcServer, func(), error) {
	wire.Build(
		ConfigSet,
		InfraSet,
		BizSet,
		RpcServerSet,
	)
	return nil, nil, nil
}
