package commands

import (
	"github.com/gogf/gf/v2/frame/g"
	log "github.com/zxdstyle/liey-admin/framework/logger"
)

var logger = g.Log().Stack(false)

func init() {
	logger.SetHandlers(log.LoggingColorHandler)
}
