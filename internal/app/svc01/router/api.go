package router

import (
	"github.com/gw-gong/go-template-project/internal/app/svc01/handler/test"
	"github.com/gw-gong/go-template-project/internal/pkg/biz/biz01"
	"github.com/gw-gong/go-template-project/internal/pkg/biz/biz02"
	"github.com/gw-gong/go-template-project/internal/pkg/client/rpc/svc02"
	"github.com/gw-gong/go-template-project/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	Biz01           biz01.Biz01
	Biz02           biz02.Biz02
	Test01DbManager mysql.Test01DbManager
	Test02DbManager mysql.Test02DbManager
	Test01Client    *svc02.Test01Client
	Test02Client    *svc02.Test02Client
}

func (r *ApiRouter) Bind(router *gin.Engine) {
	apiGroup := router.Group("/api")

	group01 := apiGroup.Group("/group01")
	{
		group01.POST("/test", test.NewTest01ApiHandler(r.Biz01))
	}

	group02 := apiGroup.Group("/group02")
	{
		group02.POST("/test", test.NewTest02ApiHandler(r.Biz02, r.Test01DbManager, r.Test02DbManager, r.Test01Client, r.Test02Client))
	}
}
