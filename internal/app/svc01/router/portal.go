package router

import (
	"github.com/gw-gong/boilerplate-go/internal/app/svc01/handler/test"
	"github.com/gw-gong/boilerplate-go/internal/pkg/biz/biz01"
	"github.com/gw-gong/boilerplate-go/internal/pkg/biz/biz02"
	"github.com/gw-gong/boilerplate-go/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

type PortalRouter struct {
	Biz01           biz01.Biz01
	Biz02           biz02.Biz02
	Test01DbManager mysql.Test01DbManager
	Test02DbManager mysql.Test02DbManager
}

func (r *PortalRouter) Bind(router *gin.Engine) {
	portalGroup := router.Group("/portal")

	group01 := portalGroup.Group("/group01")
	{
		group01.POST("/test", test.NewTest01PortalHandler(r.Biz01))
	}

	group02 := portalGroup.Group("/group02")
	{
		group02.POST("/testx", test.NewTest02PortalHandler(r.Biz02, r.Test01DbManager, r.Test02DbManager))
	}
}
