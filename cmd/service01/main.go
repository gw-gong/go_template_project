package main

import (
	"github.com/gin-gonic/gin"
	global_settings "github.com/gw-gong/gwkit-go/global_settings"
	"github.com/gw-gong/gwkit-go/log"
	utils_comm "github.com/gw-gong/gwkit-go/utils/common"
	"github.com/gw-gong/template_project/config/service01"
)

func main() {
	ctx := global_settings.GetServiceContext()
	gin.SetMode(gin.ReleaseMode)

	config, err := service01.InitConfig()
	utils_comm.ExitOnErr(ctx, err)

	syncFn, err := log.InitGlobalLogger(config.LogConfig)
	utils_comm.ExitOnErr(ctx, err)
	defer syncFn()

	httpServer, cleanup, err := InitHttpServer(
		ctx,
		config.Component01erOptions,
		config.Component02erOptions,
	)
	utils_comm.ExitOnErr(ctx, err)
	defer cleanup()
	httpServer.Run()
}
