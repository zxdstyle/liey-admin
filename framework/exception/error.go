package exception

import (
	"github.com/gogf/gf/v2/errors/gerror"
)

// kernel
var (
	ErrNotRegisterKernel      = gerror.NewCode(CodeInternalError, "please register the kernel")
	ErrRepeatedRegisterKernel = gerror.NewCode(CodeInternalError, "please do not repeated register the kernel")

	ErrInvalidOrderArgument = gerror.NewCode(CodeBadRequest, "invalid sort parameter")
)

// migration
var (
	ErrRepeatedMigration = gerror.NewCode(CodeBadRequest, "migration already exists")
)

// plugin
var (
	ErrPluginExists = gerror.NewCode(CodeInternalError, "duplicate plugin name or duplicate registered plugin")
)

// auth
var (
	ErrUnauthorized = gerror.NewCode(CodeUnauthorized, "unauthorized")
	ErrMissingToken = gerror.NewCode(CodeInternalError, "token not provider")
	ErrInvalidToken = gerror.NewCode(CodeUnauthorized, "invalid token")
)
