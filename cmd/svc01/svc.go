package main

import (
	"context"
	"fmt"

	"github.com/gw-gong/go-template-project/internal/app/svc01/router"
	"github.com/gw-gong/go-template-project/internal/config/svc01/localcfg"
	"github.com/gw-gong/go-template-project/internal/config/svc01/netcfg"

	"github.com/gin-gonic/gin"
	"github.com/gw-gong/gwkit-go/gin/middleware"
	"github.com/gw-gong/gwkit-go/log"
	"github.com/gw-gong/gwkit-go/util"
)

type HttpServer struct {
	cfg           *localcfg.Config
	netCfg        *netcfg.Config
	apiRouter     *router.ApiRouter
	appRouter     *router.AppRouter
	portalRouter  *router.PortalRouter
	privateRouter *router.PrivateRouter
}

func (s *HttpServer) Run(ctx context.Context) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	middleware.BindBasicMiddlewares(router, s.cfg.LogHttpInfo)

	s.apiRouter.Bind(router)
	s.appRouter.Bind(router)
	s.portalRouter.Bind(router)
	s.privateRouter.Bind(router)

	// start http server
	log.Infoc(ctx, "http server running", log.Int("port", s.cfg.HttpServer.Port))
	util.ExitOnErr(ctx, router.Run(fmt.Sprintf(":%d", s.cfg.HttpServer.Port)))
}
