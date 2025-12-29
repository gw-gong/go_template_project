package router

import (
	"github.com/gw-gong/boilerplate-go/internal/app/svc01/handler/test"
	"github.com/gw-gong/boilerplate-go/internal/pkg/biz/biz01"
	"github.com/gw-gong/boilerplate-go/internal/pkg/biz/biz02"
	"github.com/gw-gong/boilerplate-go/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

type PrivateRouter struct {
	Biz01           biz01.Biz01
	Biz02           biz02.Biz02
	Test01DbManager mysql.Test01DbManager
	Test02DbManager mysql.Test02DbManager
}

func (r *PrivateRouter) Bind(router *gin.Engine) {
	privateGroup := router.Group("/private")

	group01 := privateGroup.Group("/group01")
	{
		group01.POST("/test", test.NewTest01PrivateHandler(r.Biz01))
	}

	group02 := privateGroup.Group("/group02")
	{
		group02.POST("/testx", test.NewTest02PrivateHandler(r.Biz02, r.Test01DbManager, r.Test02DbManager))
	}
}
