package router

import (
	"github.com/gw-gong/go-template-project/internal/app/service01/handler/private"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component01"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component02"
	"github.com/gw-gong/go-template-project/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

type PrivateRouter struct {
	Component01  component01.Component01
	Component02  component02.Component02
	XxxDbManager mysql.XxxDbManager
}

func (r *PrivateRouter) Bind(app *gin.Engine) {
	privateGroup := app.Group("/private")

	routerGroup01 := privateGroup.Group("/router_group01")
	{
		routerGroup01.POST("/test", private.NewTestHandler(r.Component01))
	}

	routerGroup02 := privateGroup.Group("/router_group02")
	{
		routerGroup02.POST("/testx", private.NewTestxHandler(r.Component02, r.XxxDbManager))
	}
}
