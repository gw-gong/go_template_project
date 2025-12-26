//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

//go:generate go run github.com/google/wire/cmd/wire

import (
	"context"

	"github.com/google/wire"
	local_cfg "github.com/gw-gong/go-template-project/config/local_config/service01"
	"github.com/gw-gong/go-template-project/internal/app/service01/router"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component01"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component02"
	"github.com/gw-gong/go-template-project/internal/pkg/db/mysql"
)

func InitHttpServer(ctx context.Context, config *local_cfg.Config) (*httpServer, func(), error) {
	wire.Build(
		wire.FieldsOf(new(*local_cfg.Config),
			"Component01",
			"Component02"),
		component01.NewComponent01,
		component02.NewComponent02,
		mysql.NewXxxDbManager,
		wire.Struct(new(router.ApiRouter), "*"),
		wire.Struct(new(router.AppRouter), "*"),
		wire.Struct(new(router.PortalRouter), "*"),
		wire.Struct(new(router.PrivateRouter), "*"),
		wire.Struct(new(httpServer), "*"),
	)
	return nil, nil, nil
}
