package publish

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/zxdstyle/liey-admin/framework/exception"
)

var (
	ErrRepeatedPublish = gerror.NewCode(exception.CodeBadRequest, "publish already exists")
)
