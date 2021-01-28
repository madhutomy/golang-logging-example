package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
)

func SetFormatter(formatterStr string) logrus.Formatter {
	if "text" == formatterStr {
		formatter := &logrus.TextFormatter{
			ForceColors:               true,
			DisableColors:             false,
			ForceQuote:                false,
			DisableQuote:              false,
			EnvironmentOverrideColors: false,
			DisableTimestamp:          false,
			FullTimestamp:             true,
			TimestampFormat:           "2006-01-02 15:04:05",
			DisableSorting:            false,
			SortingFunc:               nil,
			DisableLevelTruncation:    false,
			PadLevelText:              false,
			QuoteEmptyFields:          false,
			FieldMap:                  nil,
			CallerPrettyfier:          nil,
		}
		return formatter
	} else {
		formatter := &logrus.JSONFormatter{
			TimestampFormat:   "01-01-2021 15:04:05",
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			FieldMap:          nil,
			CallerPrettyfier:  nil,
			PrettyPrint:       false,
		}
		return formatter
	}
}

func InitGoRoutineMaps() map[string]string {
	goRoutineMap := make(map[string]string)
	goRoutineMap["PEER_SYNC"] = "PEER_SYNC-" + strconv.Itoa(rand.Intn(200))
	goRoutineMap["PEER_STATUS"] = "PEER_STATUS-" + strconv.Itoa(rand.Intn(200))
	goRoutineMap["MONGO_SYNC"] = "MONGO_SYNC-" + strconv.Itoa(rand.Intn(200))

	return goRoutineMap
}

func CreateLogEntryForConcurrent(logger *logrus.Logger, goRoutineId string) *logrus.Entry {
	logEntry := logger.WithField("thread:", goRoutineId)
	return logEntry
}

func ChangeLogLevel(logger *logrus.Logger, logLevel string) {
	fmt.Printf(" ^^^^^^^    Log level changed, the new one is : %v ^^^^^^", logLevel)
	level, _ := logrus.ParseLevel(logLevel)
	logger.SetLevel(level)
}
