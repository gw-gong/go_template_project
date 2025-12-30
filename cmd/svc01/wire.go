//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/gw-gong/boilerplate-go/internal/app/svc01/router"
	"github.com/gw-gong/boilerplate-go/internal/config/svc01/localcfg"
	"github.com/gw-gong/boilerplate-go/internal/config/svc01/netcfg"
	"github.com/gw-gong/boilerplate-go/internal/pkg/biz/biz01"
	"github.com/gw-gong/boilerplate-go/internal/pkg/biz/biz02"
	"github.com/gw-gong/boilerplate-go/internal/pkg/client/rpc/svc02"
	"github.com/gw-gong/boilerplate-go/internal/pkg/db/mysql"
	"github.com/gw-gong/boilerplate-go/internal/pkg/util/provider"

	"github.com/google/wire"
	"github.com/gw-gong/gwkit-go/grpc/consul"
	"github.com/gw-gong/gwkit-go/hotcfg"
)

var ConfigSet = wire.NewSet(
	localcfg.NewConfig,
	wire.FieldsOf(
		new(*localcfg.Config),
		"ConsulAgentAddr",
		"ConsulNetCfg",
		"Biz01",
		"Biz02",
		"Test01Client",
		"Test02Client",
	),
	netcfg.NewConfig,
	wire.FieldsOf(
		new(*netcfg.Config),
		"Test01DbManager",
		"Test02DbManager",
	),
	hotcfg.NewHotLoaderManager,
)

var InfraSet = wire.NewSet(
	provider.NewGinEngine,
	mysql.NewTest01DbManager,
	mysql.NewTest02DbManager,
	consul.NewConsulClient,
	svc02.NewTest01Client,
	svc02.NewTest02Client,
)

var BizSet = wire.NewSet(
	biz01.NewBiz01,
	biz02.NewBiz02,
)

var RouterSet = wire.NewSet(
	wire.Struct(new(router.ApiRouter), "*"),
	wire.Struct(new(router.AppRouter), "*"),
	wire.Struct(new(router.PortalRouter), "*"),
	wire.Struct(new(router.PrivateRouter), "*"),
)

var HttpServerSet = wire.NewSet(
	wire.Struct(new(HttpServer), "*"),
)

func InitHttpServer(cfgOption *hotcfg.LocalConfigOption) (*HttpServer, func(), error) {
	wire.Build(
		ConfigSet,
		InfraSet,
		BizSet,
		RouterSet,
		HttpServerSet,
	)
	return nil, nil, nil
}
