package logger

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"os"
)

type JsonOutputsForLogger struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Content string `json:"content"`
}

var (
	LoggingJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
		jsonForLogger := JsonOutputsForLogger{
			Time:    in.TimeFormat,
			Level:   gstr.Trim(in.LevelFormat, "[]"),
			Content: gstr.Trim(in.Content),
		}
		jsonBytes, err := json.Marshal(jsonForLogger)
		if err != nil {
			_, _ = os.Stderr.WriteString(err.Error())
			return
		}

		in.Buffer.WriteString(colorFormat(in.Level, string(jsonBytes)))
		in.Buffer.WriteString("\n")
		in.Next(ctx)
	}
)
