package router

import (
	"github.com/gw-gong/go-template-project/internal/app/svc01/handlertest"
	"github.com/gw-gong/go-template-project/internal/pkg/biz/biz01"
	"github.com/gw-gong/go-template-project/internal/pkg/biz/biz02"
	"github.com/gw-gong/go-template-project/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	Biz01           biz01.Biz01
	Biz02           biz02.Biz02
	Test01DbManager mysql.Test01DbManager
	Test02DbManager mysql.Test02DbManager
}

func (r *AppRouter) Bind(router *gin.Engine) {
	appGroup := router.Group("/app")

	group01 := appGroup.Group("/group01")
	{
		group01.POST("/test", handlertest.NewTest01AppHandler(r.Biz01))
	}

	group02 := appGroup.Group("/group02")
	{
		group02.POST("/testx", handlertest.NewTest02AppHandler(r.Biz02, r.Test01DbManager, r.Test02DbManager))
	}
}
