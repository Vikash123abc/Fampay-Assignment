package config

import (
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitializeLogger() {

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	customFormatter.DisableLevelTruncation = true
	customFormatter.PadLevelText = true
	customFormatter.CallerPrettyfier = func(frame *runtime.Frame) (function string, file string) {
		fileName := " | " + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line) + " | "
		return path.Base(frame.Function) + " | ", fileName
	}

	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(customFormatter)
	Logger.SetReportCaller(true)
}
