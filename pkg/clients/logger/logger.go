/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package logger

import (
	"os"
	"time"

	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
)

var appLogger *log.MyLogger

func init() {
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(getLogLevel(constants.DEFAULT_LOG_LEVEL))
	log.Logger().SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		DisableColors:   false,
	})
	// log.Logger().SetReportCaller(true)
	appLogger = log.Logger()
}
func SetLogger(lvl string) {
	appLogger.SetLevel(getLogLevel(lvl))
}

func SetDevMode() {
	SetLogger(constants.DEFAULT_DEV_LOG_LEVEL)
}

func GetLogger() *log.MyLogger {
	return appLogger
}

func Debug(msg string, args ...interface{}) {
	appLogger.Logf(logrus.DebugLevel, msg, args...)
}

func Info(msg string, args ...interface{}) {
	appLogger.Logf(logrus.InfoLevel, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	appLogger.Logf(logrus.WarnLevel, msg, args...)
}

func Error(msg string, args ...interface{}) {
	appLogger.Logf(logrus.ErrorLevel, msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	appLogger.Logf(logrus.FatalLevel, msg, args...)
}

func Panic(msg string, args ...interface{}) {
	appLogger.Logf(logrus.PanicLevel, msg, args...)
}
