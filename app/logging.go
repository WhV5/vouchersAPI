/**
* @Author : henry
* @Data: 2020-08-13 13:24
* @Note:
**/

package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

var Logger *logrus.Logger

type appLoggerHook struct {
}

func (appLoggerHook *appLoggerHook) Fire(entry *logrus.Entry) error {
	if entry.Level == logrus.ErrorLevel || entry.Level == logrus.FatalLevel {
		println("STOP_PIPELINES")
	}
	return nil
}

func (appLoggerHook *appLoggerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func init() {
	Logger = logrus.New()
	Logger.SetReportCaller(true)
	Logger.AddHook(&appLoggerHook{})
	Logger.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := strings.Join(strings.Split(f.File, "/vouchersAPI")[1:], "/vouchersAPI")
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	initConfig()
}
