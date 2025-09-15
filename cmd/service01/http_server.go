package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gw-gong/gwkit-go/gin/middlewares"
	utils_comm "github.com/gw-gong/gwkit-go/utils/common"
	"github.com/gw-gong/template_project/internal/app/service01/router"
)

type httpServer struct {
	ctx       context.Context
	apiRouter *router.ApiRouter
}

func (s *httpServer) Run() {
	app := gin.New()
	app.Use(middlewares.Cors(
		middlewares.WithCorsOptAllowOrigins([]string{"https://aisuda.github.io"}),
	))
	middlewares.BindBasicMiddlewares(app, true)

	s.apiRouter.Bind(app)

	utils_comm.ExitOnErr(s.ctx, app.Run(":8080"))
}
