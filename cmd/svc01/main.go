package main

import (
	"github.com/gw-gong/go-template-project/internal/config/svc01/localcfg"
	"github.com/gw-gong/go-template-project/internal/config/svc01/netcfg"

	"github.com/gw-gong/gwkit-go/hotcfg"
	"github.com/gw-gong/gwkit-go/log"
	"github.com/gw-gong/gwkit-go/setting"
	"github.com/gw-gong/gwkit-go/util"
)

func main() {
	ctx := setting.GetServiceContext()
	hlm := hotcfg.NewHotLoaderManager()

	cfgPath, cfgFileName, err := initFlags()
	util.ExitOnErr(ctx, err)

	// load local config
	localCfg, err := localcfg.NewConfig(&hotcfg.LocalConfigOption{
		FilePath: cfgPath,
		FileName: cfgFileName,
		FileType: "yaml",
	})
	util.ExitOnErr(ctx, err)
	util.ExitOnErr(ctx, hlm.RegisterHotLoader(localCfg))

	// set service env
	setting.SetEnv(localCfg.Env)

	// init global logger
	syncFn, err := log.InitGlobalLogger(localCfg.Logger)
	util.ExitOnErr(ctx, err)
	defer syncFn()

	// load net config
	netCfg, err := netcfg.NewConfig(localCfg.ConsulNetCfg)
	util.ExitOnErr(ctx, err)
	util.ExitOnErr(ctx, hlm.RegisterHotLoader(netCfg))

	// start hot reload
	util.ExitOnErr(ctx, hlm.Watch())

	// init & run http server
	httpServer, close, err := InitHttpServer(localCfg, netCfg)
	util.ExitOnErr(ctx, err)
	defer close()
	httpServer.Run(ctx)
}
