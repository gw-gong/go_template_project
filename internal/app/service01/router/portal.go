package router

import (
	"github.com/gw-gong/go-template-project/internal/app/service01/handler/portal"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component01"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component02"
	"github.com/gw-gong/go-template-project/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

type PortalRouter struct {
	Component01  component01.Component01
	Component02  component02.Component02
	XxxDbManager mysql.XxxDbManager
}

func (r *PortalRouter) Bind(app *gin.Engine) {
	portalGroup := app.Group("/portal")

	routerGroup01 := portalGroup.Group("/router_group01")
	{
		routerGroup01.POST("/test", portal.NewTestHandler(r.Component01))
	}

	routerGroup02 := portalGroup.Group("/router_group02")
	{
		routerGroup02.POST("/testx", portal.NewTestxHandler(r.Component02, r.XxxDbManager))
	}
}
