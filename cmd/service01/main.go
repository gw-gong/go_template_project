package main

import (
	"flag"

	local_cfg "github.com/gw-gong/go-template-project/config/local_config/service01"

	"github.com/gin-gonic/gin"
	global_settings "github.com/gw-gong/gwkit-go/global_settings"
	"github.com/gw-gong/gwkit-go/hot_cfg"
	"github.com/gw-gong/gwkit-go/log"
	utils_comm "github.com/gw-gong/gwkit-go/utils/common"
)

func main() {
	ctx := global_settings.GetServiceContext()
	gin.SetMode(gin.ReleaseMode)
	flag.Parse()
	validateFlags(ctx)

	hucm := hot_cfg.GetHotUpdateManager()
	utils_comm.ExitOnErr(ctx, local_cfg.InitConfig(*cfgFilePath, *cfgFileName, "yaml"))
	utils_comm.ExitOnErr(ctx, hucm.RegisterHotUpdateConfig(local_cfg.Cfg))
	utils_comm.ExitOnErr(ctx, hucm.Watch())

	syncFn, err := log.InitGlobalLogger(local_cfg.Cfg.Logger)
	utils_comm.ExitOnErr(ctx, err)
	defer syncFn()

	httpServer, cleanup, err := InitHttpServer(ctx, local_cfg.Cfg)
	utils_comm.ExitOnErr(ctx, err)
	defer cleanup()
	httpServer.Run()
}
