package logger

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	defaultLevelColor = map[int]func(format string, a ...interface{}) string{
		glog.LEVEL_DEBU: color.CyanString,
		glog.LEVEL_ERRO: color.HiRedString,
		glog.LEVEL_NOTI: color.HiGreenString,
		glog.LEVEL_WARN: color.HiYellowString,
		glog.LEVEL_CRIT: color.RedString,
		glog.LEVEL_PANI: color.RedString,
		glog.LEVEL_FATA: color.RedString,
	}

	LoggingColorHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
		in.Buffer.Reset()
		content := fmt.Sprintf("%s %s %s", in.TimeFormat, in.LevelFormat, in.Content)
		in.Buffer.WriteString(colorFormat(in.Level, content))
		in.Buffer.WriteString("\n")
		in.Next(ctx)
	}
)

func colorFormat(level int, str string) string {
	if colorFunc, ok := defaultLevelColor[level]; ok {
		return colorFunc(str)
	}
	return str
}
