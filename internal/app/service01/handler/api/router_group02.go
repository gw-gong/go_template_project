package api

import (
	"github.com/gw-gong/go-template-project/api/http/service01/api"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component02"
	"github.com/gw-gong/go-template-project/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
	gwkit_res "github.com/gw-gong/gwkit-go/gin/response"
	gwkit_res_code "github.com/gw-gong/gwkit-go/http/response"
)

type TestxHandler struct {
	request      *api.TestxRequest
	response     *api.TestxResponse
	component02  component02.Component02
	xxxDbManager mysql.XxxDbManager
}

func NewTestxHandler(component02 component02.Component02, xxxDbManager mysql.XxxDbManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := &TestxHandler{
			request:      &api.TestxRequest{},
			response:     &api.TestxResponse{},
			component02:  component02,
			xxxDbManager: xxxDbManager,
		}
		if err := c.ShouldBindJSON(handler.request); err != nil {
			gwkit_res.ResponseError(c, gwkit_res_code.ErrParam)
			return
		}
		handler.Handle(c)
	}
}

func (h *TestxHandler) Handle(c *gin.Context) {
	h.component02.Function02(c.Request.Context())
}
