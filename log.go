package swissknife

import (
	"github.com/sirupsen/logrus"
	"time"
)

// NewJSONLogger
// Create new logrus JSON logger with 3 FieldKey {ts,severity,msg}
// and timestamp compliant with RFC3339
func NewJSONLogger() *logrus.Logger {
	logger := logrus.Logger{}

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})
	return &logger
}
