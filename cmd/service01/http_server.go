package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	local_cfg "github.com/gw-gong/go-template-project/config/local_config/service01"
	"github.com/gw-gong/go-template-project/internal/app/service01/router"
	"github.com/gw-gong/gwkit-go/gin/middlewares"
	"github.com/gw-gong/gwkit-go/log"
	utils_comm "github.com/gw-gong/gwkit-go/utils/common"
)

type httpServer struct {
	ctx           context.Context
	cfg           *local_cfg.Config
	apiRouter     *router.ApiRouter
	appRouter     *router.AppRouter
	portalRouter  *router.PortalRouter
	privateRouter *router.PrivateRouter
}

func (s *httpServer) Run() {
	app := gin.New()
	app.Use(middlewares.Cors(
		middlewares.WithCorsOptAllowOrigins(s.cfg.HttpServer.Cors.AllowOrigins),
		middlewares.WithCorsOptAllowCredentials(s.cfg.HttpServer.Cors.AllowCredentials),
	))
	middlewares.BindBasicMiddlewares(app, true)

	s.apiRouter.Bind(app)
	s.appRouter.Bind(app)
	s.portalRouter.Bind(app)
	s.privateRouter.Bind(app)

	log.Infoc(s.ctx, "http server running", log.Int("port", s.cfg.HttpServer.Port))
	utils_comm.ExitOnErr(s.ctx, app.Run(fmt.Sprintf(":%d", s.cfg.HttpServer.Port)))
}
