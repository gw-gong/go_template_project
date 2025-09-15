//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

//go:generate go run github.com/google/wire/cmd/wire

import (
	"context"

	"github.com/google/wire"
	"github.com/gw-gong/template_project/internal/app/service01/router"
	"github.com/gw-gong/template_project/internal/pkg/components/component01"
	"github.com/gw-gong/template_project/internal/pkg/components/component02"
	"github.com/gw-gong/template_project/internal/pkg/db/mysql"
)

func InitHttpServer(
	ctx context.Context,
	component01erOptions component01.Component01erOptions,
	component02erOptions component02.Component02erOptions,
) (*httpServer, func(), error) {
	wire.Build(
		component01.NewComponent01er,
		component02.NewComponent02er,
		mysql.NewXxxDbManager,
		wire.Struct(new(router.ApiRouter), "*"),
		wire.Struct(new(httpServer), "*"),
	)
	return nil, nil, nil
}
