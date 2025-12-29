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

	"github.com/google/wire"
	"github.com/gw-gong/gwkit-go/grpc/consul"
)

func InitHttpServer(config *localcfg.Config, netCfg *netcfg.Config) (*HttpServer, func(), error) {
	wire.Build(
		wire.FieldsOf(
			new(*localcfg.Config),
			"ConsulAgentAddr",
			"Biz01",
			"Biz02",
			"Test01Client",
			"Test02Client",
		),
		wire.FieldsOf(
			new(*netcfg.Config),
			"Test01DbManager",
			"Test02DbManager",
		),
		consul.NewConsulClient,
		biz01.NewBiz01,
		biz02.NewBiz02,
		mysql.NewTest01DbManager,
		mysql.NewTest02DbManager,
		svc02.NewTest01Client,
		svc02.NewTest02Client,
		wire.Struct(new(router.ApiRouter), "*"),
		wire.Struct(new(router.AppRouter), "*"),
		wire.Struct(new(router.PortalRouter), "*"),
		wire.Struct(new(router.PrivateRouter), "*"),
		wire.Struct(new(HttpServer), "*"),
	)
	return nil, nil, nil
}
