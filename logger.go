package galaxylib

import (
	"fmt"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

type GalaxyLog struct {
	Logger *logrus.Logger
}

var DefaultGalaxyLog = &GalaxyLog{}

var GalaxyLogger *logrus.Logger

//var GalaxyLog *logrus.Logger

func (g *GalaxyLog) ConfigLogger() *logrus.Logger {
	if GalaxyLogger != nil {
		return GalaxyLogger
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./log/info/info.log",
		logrus.ErrorLevel: "./log/error/error.log",
		logrus.DebugLevel: "./log/debug/debug.log",
		logrus.FatalLevel: "./log/fatal/fatal.log",
		logrus.WarnLevel:  "./log/warn/warn.log",
	}

	//infoWriter := buildWriter()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	logrus.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  buildWriter(pathMap[logrus.InfoLevel]),
			logrus.ErrorLevel: buildWriter(pathMap[logrus.ErrorLevel]),
			logrus.WarnLevel:  buildWriter(pathMap[logrus.WarnLevel]),
		},
		&logrus.JSONFormatter{},
	))

	GalaxyLogger = logrus.New()
	GalaxyLogger.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return GalaxyLogger
}

func buildWriter(path string) *rotatelogs.RotateLogs {
	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	if err != nil {
		fmt.Println(err)
	}
	return writer
}
