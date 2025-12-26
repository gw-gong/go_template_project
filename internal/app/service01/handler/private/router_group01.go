package private

import (
	"github.com/gw-gong/go-template-project/api/http/service01/private"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component01"

	"github.com/gin-gonic/gin"
	gwkit_res "github.com/gw-gong/gwkit-go/gin/response"
	gwkit_res_code "github.com/gw-gong/gwkit-go/http/response"
	"github.com/gw-gong/gwkit-go/log"
)

type TestHandler struct {
	request     *private.TestRequest
	response    *private.TestResponse
	component01 component01.Component01
}

func NewTestHandler(component01 component01.Component01) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := &TestHandler{
			request:     &private.TestRequest{},
			response:    &private.TestResponse{},
			component01: component01,
		}
		if err := c.ShouldBindJSON(handler.request); err != nil {
			log.Errorc(c.Request.Context(), "bind json failed", log.Err(err))
			gwkit_res.ResponseErrorWithDetails(c, gwkit_res_code.ErrParam, err.Error())
			return
		}
		handler.Handle(c)
	}
}

func (h *TestHandler) Handle(c *gin.Context) {
	h.component01.Function01(c.Request.Context())

	h.response.Field01 = "test"
	gwkit_res.ResponseSuccess(c, h.response)
}
