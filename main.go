package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/WALL-EEEEEEE/gagdets/core"
	_ "github.com/WALL-EEEEEEE/gagdets/pipes"
	_ "github.com/WALL-EEEEEEE/gagdets/sources"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func main() {
	log.SetReportCaller(true)
	log.SetFormatter(&zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
		Formatter: nested.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			ShowFullLevel:   true,
		},
	})
	core.Exec.Start()
}
