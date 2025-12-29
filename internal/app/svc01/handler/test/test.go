package test

import (
	"context"

	"github.com/gw-gong/boilerplate-go/api/http/svc01/api"
	"github.com/gw-gong/boilerplate-go/api/http/svc01/app"
	"github.com/gw-gong/boilerplate-go/api/http/svc01/portal"
	"github.com/gw-gong/boilerplate-go/api/http/svc01/private"
	"github.com/gw-gong/boilerplate-go/internal/app/svc01/errcode"
	"github.com/gw-gong/boilerplate-go/internal/pkg/biz/biz01"
	"github.com/gw-gong/boilerplate-go/internal/pkg/biz/biz02"
	"github.com/gw-gong/boilerplate-go/internal/pkg/client/rpc/svc02"
	"github.com/gw-gong/boilerplate-go/internal/pkg/db/mysql"

	"github.com/gin-gonic/gin"
	"github.com/gw-gong/gwkit-go/gin/res"
	"github.com/gw-gong/gwkit-go/gin/validator"
	"github.com/gw-gong/gwkit-go/log"
)

type Test01ApiHandler struct {
	ctx      context.Context
	request  *api.TestRequest01
	response *api.TestResponse01
	Biz01    biz01.Biz01
}

func NewTest01ApiHandler(biz01 biz01.Biz01) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := &Test01ApiHandler{
			ctx:      c.Request.Context(),
			request:  &api.TestRequest01{},
			response: &api.TestResponse01{},
			Biz01:    biz01,
		}
		if err := c.ShouldBindJSON(handler.request); err != nil {
			fmtErr := validator.FmtValidationErrors(err, handler.request)
			res.ResponseError(c, errcode.ErrCodeRequestParamInvalid, res.WithErrDetail(fmtErr))
			return
		}
		handler.handle(c)
	}
}

func (h *Test01ApiHandler) handle(c *gin.Context) {
	log.Infoc(c.Request.Context(), "Test01ApiHandler handle")
	h.Biz01.Function01(h.ctx)
	res.ResponseSuccess(c, nil)
}

type Test02ApiHandler struct {
	ctx             context.Context
	request         *api.TestRequest02
	response        *api.TestResponse02
	Biz02           biz02.Biz02
	Test01DbManager mysql.Test01DbManager
	Test02DbManager mysql.Test02DbManager
	Test01Client    *svc02.Test01Client
	Test02Client    *svc02.Test02Client
}

func NewTest02ApiHandler(biz02 biz02.Biz02, test01DbManager mysql.Test01DbManager, test02DbManager mysql.Test02DbManager, test01Client *svc02.Test01Client, test02Client *svc02.Test02Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := &Test02ApiHandler{
			ctx:             c.Request.Context(),
			request:         &api.TestRequest02{},
			response:        &api.TestResponse02{},
			Biz02:           biz02,
			Test01DbManager: test01DbManager,
			Test02DbManager: test02DbManager,
			Test01Client:    test01Client,
			Test02Client:    test02Client,
		}
		if err := c.ShouldBindJSON(handler.request); err != nil {
			fmtErr := validator.FmtValidationErrors(err, handler.request)
			res.ResponseError(c, errcode.ErrCodeRequestParamInvalid, res.WithErrDetail(fmtErr))
			return
		}
		handler.handle(c)
	}
}

func (h *Test02ApiHandler) handle(c *gin.Context) {
	log.Infoc(c.Request.Context(), "Test02ApiHandler handle")
	h.Biz02.Function01(h.ctx)
	h.Test01DbManager.Setxxxx(h.ctx)
	h.Test02DbManager.Setxxxx(h.ctx)
	resField01, resField02, err := h.Test01Client.TestFunc(h.ctx, h.request.Field01, h.request.Field02)
	if err != nil {
		res.ResponseError(c, errcode.ErrCodeInternalServerError, res.WithErrDetail(err))
		return
	}
	log.Infoc(h.ctx, "Test01Client.TestFunc", log.Str("resField01", resField01), log.Str("resField02", resField02))
	resField01, resField02, err = h.Test02Client.TestFunc(h.ctx, h.request.Field01, h.request.Field02)
	if err != nil {
		res.ResponseError(c, errcode.ErrCodeInternalServerError, res.WithErrDetail(err))
		return
	}
	log.Infoc(h.ctx, "Test02Client.TestFunc", log.Str("resField01", resField01), log.Str("resField02", resField02))
	res.ResponseSuccess(c, nil)
}

type Test01AppHandler struct {
	ctx      context.Context
	request  *app.TestRequest01
	response *app.TestResponse01
	Biz01    biz01.Biz01
}

func NewTest01AppHandler(biz01 biz01.Biz01) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := &Test01AppHandler{
			ctx:      c.Request.Context(),
			request:  &app.TestRequest01{},
			response: &app.TestResponse01{},
			Biz01:    biz01,
		}
		if err := c.ShouldBindJSON(handler.request); err != nil {
			fmtErr := validator.FmtValidationErrors(err, handler.request)
			res.ResponseError(c, errcode.ErrCodeRequestParamInvalid, res.WithErrDetail(fmtErr))
			return
		}
		handler.handle(c)
	}
}

func (h *Test01AppHandler) handle(c *gin.Context) {
	log.Infoc(c.Request.Context(), "Test01AppHandler handle")
	h.Biz01.Function01(h.ctx)
	res.ResponseSuccess(c, nil)
}

type Test02AppHandler struct {
	ctx             context.Context
	request         *app.TestRequest02
	response        *app.TestResponse02
	Biz02           biz02.Biz02
	Test01DbManager mysql.Test01DbManager
	Test02DbManager mysql.Test02DbManager
}

func NewTest02AppHandler(biz02 biz02.Biz02, test01DbManager mysql.Test01DbManager, test02DbManager mysql.Test02DbManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := &Test02AppHandler{
			ctx:             c.Request.Context(),
			request:         &app.TestRequest02{},
			response:        &app.TestResponse02{},
			Biz02:           biz02,
			Test01DbManager: test01DbManager,
			Test02DbManager: test02DbManager,
		}
		if err := c.ShouldBindJSON(handler.request); err != nil {
			fmtErr := validator.FmtValidationErrors(err, handler.request)
			res.ResponseError(c, errcode.ErrCodeRequestParamInvalid, res.WithErrDetail(fmtErr))
			return
		}
		handler.handle(c)
	}
}

func (h *Test02AppHandler) handle(c *gin.Context) {
	log.Infoc(c.Request.Context(), "Test02AppHandler handle")
	h.Biz02.Function01(h.ctx)
	h.Test01DbManager.Setxxxx(h.ctx)
	h.Test02DbManager.Setxxxx(h.ctx)
	res.ResponseSuccess(c, nil)
}

type Test01PortalHandler struct {
	ctx      context.Context
	request  *portal.TestRequest01
	response *portal.TestResponse01
	Biz01    biz01.Biz01
}

func NewTest01PortalHandler(biz01 biz01.Biz01) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := &Test01PortalHandler{
			ctx:      c.Request.Context(),
			request:  &portal.TestRequest01{},
			response: &portal.TestResponse01{},
			Biz01:    biz01,
		}
		if err := c.ShouldBindJSON(handler.request); err != nil {
			fmtErr := validator.FmtValidationErrors(err, handler.request)
			res.ResponseError(c, errcode.ErrCodeRequestParamInvalid, res.WithErrDetail(fmtErr))
			return
		}
		handler.handle(c)
	}
}

func (h *Test01PortalHandler) handle(c *gin.Context) {
	log.Infoc(c.Request.Context(), "Test01PortalHandler handle")
	h.Biz01.Function01(h.ctx)
	res.ResponseSuccess(c, nil)
}

type Test02PortalHandler struct {
	ctx             context.Context
	request         *portal.TestRequest02
	response        *portal.TestResponse02
	Biz02           biz02.Biz02
	Test01DbManager mysql.Test01DbManager
	Test02DbManager mysql.Test02DbManager
}

func NewTest02PortalHandler(biz02 biz02.Biz02, test01DbManager mysql.Test01DbManager, test02DbManager mysql.Test02DbManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := &Test02PortalHandler{
			ctx:             c.Request.Context(),
			request:         &portal.TestRequest02{},
			response:        &portal.TestResponse02{},
			Biz02:           biz02,
			Test01DbManager: test01DbManager,
			Test02DbManager: test02DbManager,
		}
		if err := c.ShouldBindJSON(handler.request); err != nil {
			fmtErr := validator.FmtValidationErrors(err, handler.request)
			res.ResponseError(c, errcode.ErrCodeRequestParamInvalid, res.WithErrDetail(fmtErr))
			return
		}
		handler.handle(c)
	}
}

func (h *Test02PortalHandler) handle(c *gin.Context) {
	log.Infoc(c.Request.Context(), "Test02PortalHandler handle")
	h.Biz02.Function01(h.ctx)
	h.Test01DbManager.Setxxxx(h.ctx)
	h.Test02DbManager.Setxxxx(h.ctx)
	res.ResponseSuccess(c, nil)
}

type Test01PrivateHandler struct {
	ctx      context.Context
	request  *private.TestRequest01
	response *private.TestResponse01
	Biz01    biz01.Biz01
}

func NewTest01PrivateHandler(biz01 biz01.Biz01) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := &Test01PrivateHandler{
			ctx:      c.Request.Context(),
			request:  &private.TestRequest01{},
			response: &private.TestResponse01{},
			Biz01:    biz01,
		}
		if err := c.ShouldBindJSON(handler.request); err != nil {
			fmtErr := validator.FmtValidationErrors(err, handler.request)
			res.ResponseError(c, errcode.ErrCodeRequestParamInvalid, res.WithErrDetail(fmtErr))
			return
		}
		handler.handle(c)
	}
}

func (h *Test01PrivateHandler) handle(c *gin.Context) {
	log.Infoc(c.Request.Context(), "Test01PrivateHandler handle")
	h.Biz01.Function01(h.ctx)
	res.ResponseSuccess(c, nil)
}

type Test02PrivateHandler struct {
	ctx             context.Context
	request         *private.TestRequest02
	response        *private.TestResponse02
	Biz02           biz02.Biz02
	Test01DbManager mysql.Test01DbManager
	Test02DbManager mysql.Test02DbManager
}

func NewTest02PrivateHandler(biz02 biz02.Biz02, test01DbManager mysql.Test01DbManager, test02DbManager mysql.Test02DbManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := &Test02PrivateHandler{
			ctx:             c.Request.Context(),
			request:         &private.TestRequest02{},
			response:        &private.TestResponse02{},
			Biz02:           biz02,
			Test01DbManager: test01DbManager,
			Test02DbManager: test02DbManager,
		}
		if err := c.ShouldBindJSON(handler.request); err != nil {
			fmtErr := validator.FmtValidationErrors(err, handler.request)
			res.ResponseError(c, errcode.ErrCodeRequestParamInvalid, res.WithErrDetail(fmtErr))
			return
		}
		handler.handle(c)
	}
}

func (h *Test02PrivateHandler) handle(c *gin.Context) {
	log.Infoc(c.Request.Context(), "Test02PrivateHandler handle")
	h.Biz02.Function01(h.ctx)
	h.Test01DbManager.Setxxxx(h.ctx)
	h.Test02DbManager.Setxxxx(h.ctx)
	res.ResponseSuccess(c, nil)
}
