package router

import (
	app_handler "github.com/gw-gong/go-template-project/internal/app/service01/handler/app"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component01"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component02"
	"github.com/gw-gong/go-template-project/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	Component01  component01.Component01
	Component02  component02.Component02
	XxxDbManager mysql.XxxDbManager
}

func (r *AppRouter) Bind(app *gin.Engine) {
	appGroup := app.Group("/app")

	routerGroup01 := appGroup.Group("/router_group01")
	{
		routerGroup01.POST("/test", app_handler.NewTestHandler(r.Component01))
	}

	routerGroup02 := appGroup.Group("/router_group02")
	{
		routerGroup02.POST("/testx", app_handler.NewTestxHandler(r.Component02, r.XxxDbManager))
	}
}
