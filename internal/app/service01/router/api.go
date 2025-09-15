package router

import (
	"github.com/gw-gong/template_project/internal/app/service01/handler/api"
	"github.com/gw-gong/template_project/internal/pkg/components/component01"
	"github.com/gw-gong/template_project/internal/pkg/components/component02"
	"github.com/gw-gong/template_project/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	Component01er component01.Component01er
	Component02er component02.Component02er
	XxxDbManager  mysql.XxxDbManager
}

func (r *ApiRouter) Bind(app *gin.Engine) {
	apiGroup := app.Group("/api")

	routerGroup01 := apiGroup.Group("/router_group01")
	{
		routerGroup01.POST("/test", api.NewTestHandler(r.Component01er))
	}

	routerGroup02 := apiGroup.Group("/router_group02")
	{
		routerGroup02.POST("/testx", api.NewTestxHandler(r.Component02er, r.XxxDbManager))
	}
}
