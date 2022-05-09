package jwt

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/zxdstyle/liey-admin/framework/exception"
)

var (
	TokenExpired     = gerror.NewCode(exception.CodeUnauthorized, "token is expired")
	TokenNotValidYet = gerror.NewCode(exception.CodeUnauthorized, "token not active yet")
	TokenMalformed   = gerror.NewCode(exception.CodeUnauthorized, "that's not even a token")
	TokenInvalid     = gerror.NewCode(exception.CodeUnauthorized, "couldn't handle this token")
)
