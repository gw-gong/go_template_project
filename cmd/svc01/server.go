package main

import (
	"context"
	"fmt"

	"github.com/gw-gong/boilerplate-go/internal/app/svc01/router"
	"github.com/gw-gong/boilerplate-go/internal/config/svc01/localcfg"
	"github.com/gw-gong/boilerplate-go/internal/config/svc01/netcfg"

	"github.com/gin-gonic/gin"
	"github.com/gw-gong/gwkit-go/gin/middleware"
	"github.com/gw-gong/gwkit-go/hotcfg"
	"github.com/gw-gong/gwkit-go/log"
	"github.com/gw-gong/gwkit-go/setting"
	"github.com/gw-gong/gwkit-go/util"
)

type HttpServer struct {
	localCfg      *localcfg.Config
	netCfg        *netcfg.Config
	hlm           hotcfg.HotLoaderManager
	router        *gin.Engine
	apiRouter     *router.ApiRouter
	appRouter     *router.AppRouter
	portalRouter  *router.PortalRouter
	privateRouter *router.PrivateRouter
}

func (s *HttpServer) SetupAndRun(ctx context.Context) {
	setting.SetEnv(s.localCfg.Env)

	syncFn, err := log.InitGlobalLogger(s.localCfg.Logger)
	util.ExitOnErr(ctx, err)

	util.ExitOnErr(ctx, s.hlm.RegisterHotLoader(s.localCfg))
	util.ExitOnErr(ctx, s.hlm.RegisterHotLoader(s.netCfg))
	util.ExitOnErr(ctx, s.hlm.Watch())

	middleware.BindBasicMiddlewares(s.router, s.localCfg.LogHttpInfo)

	s.apiRouter.Bind(s.router)
	s.appRouter.Bind(s.router)
	s.portalRouter.Bind(s.router)
	s.privateRouter.Bind(s.router)

	log.Infoc(ctx, "http server running", log.Int("port", s.localCfg.HttpServer.Port))
	util.ExitOnErr(ctx, s.router.Run(fmt.Sprintf(":%d", s.localCfg.HttpServer.Port)))
	syncFn()
}
