package swissknife

import (
	"github.com/sirupsen/logrus"
	"time"
	"os"
)

// NewJSONLogger
// Create new logrus JSON logger with 3 FieldKey {ts,severity,msg}
// and timestamp compliant with RFC3339
func NewJSONLogger() *logrus.Logger {
	logger := &logrus.Logger{
		Out: os.Stderr,
		Hooks: make(logrus.LevelHooks),
		Level: logrus.DebugLevel,
	}

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})
	return logger
}
