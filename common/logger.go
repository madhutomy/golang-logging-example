package common

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

var lock = &sync.Mutex{}
var logger *logrus.Logger
const loggerID = "loggerID"

func GetLogger() *logrus.Logger {
	lock.Lock()
	defer lock.Unlock()

	if logger == nil {
		logger = logrus.New()
		logger.SetLevel(logrus.InfoLevel)
		logger.SetFormatter(SetFormatter("text"))
		logger.SetReportCaller(true)
	}

	return logger
}

func GetLoggerWithContext(ctx context.Context) *logrus.Entry {
	var id interface{}

	logger = GetLogger()
	if ctx != nil && ctx.Value(loggerID) != nil {
		id = ctx.Value(loggerID)
		if id != nil {
			logEntry := logger.WithField("[routine id]", id)
			return logEntry
		}
	}

	return logrus.NewEntry(logger)
}

func ModifyLogLevel(logLevel string) *logrus.Logger {
	level, _ := logrus.ParseLevel(logLevel)
	logger = GetLogger()
	logger.SetLevel(level)
	logger.Printf("Log level changed, the new one is : ", logLevel)

	return logger
}

func CtxWithLoggerID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, loggerID, id)
}