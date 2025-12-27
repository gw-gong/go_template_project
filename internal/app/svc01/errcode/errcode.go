package errcode

import (
	"net/http"

	"github.com/gw-gong/gwkit-go/http/code"
)

var ErrCodeSuccess = code.NewErrCode(0, "success", http.StatusOK)

// client error
var (
	ErrCodeRequestParamInvalid = code.NewErrCode(100000000, "request param invalid", http.StatusOK)
)

// server error
var (
	ErrCodeInternalServerError = code.NewErrCode(200000000, "internal server error", http.StatusOK)
)
