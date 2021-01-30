package common

import (
	"github.com/sirupsen/logrus"
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
			TimestampFormat:           "2006-01-02T15:04:05Z-07:00",
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
			TimestampFormat:   "2006-01-02T15:04:05Z-07:00",
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
