package router

import (
	"github.com/gw-gong/go-template-project/internal/app/service01/handler/api"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component01"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component02"
	"github.com/gw-gong/go-template-project/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	Component01  component01.Component01
	Component02  component02.Component02
	XxxDbManager mysql.XxxDbManager
}

func (r *ApiRouter) Bind(app *gin.Engine) {
	apiGroup := app.Group("/api")

	routerGroup01 := apiGroup.Group("/router_group01")
	{
		routerGroup01.POST("/test", api.NewTestHandler(r.Component01))
	}

	routerGroup02 := apiGroup.Group("/router_group02")
	{
		routerGroup02.POST("/testx", api.NewTestxHandler(r.Component02, r.XxxDbManager))
	}
}
