package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gw-gong/gwkit-go/setting"
	"github.com/gw-gong/gwkit-go/util"
)

func main() {
	ctx := setting.GetServiceContext()
	gin.SetMode(gin.ReleaseMode)

	cfgOption, err := initFlags()
	util.ExitOnErr(ctx, err)

	server, cleanup, err := InitHttpServer(cfgOption)
	util.ExitOnErr(ctx, err)
	defer cleanup()

	server.SetupAndRun(ctx)
}
