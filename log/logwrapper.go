package logwrapper

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// log level
// denote the log is info, warning, error, etc.
type EventType int

const (
	Info EventType = iota
	Warning
	Error
)

// Event stores messages to log later, from our standard interface
type Event struct {
	eventType EventType
	message   string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Necessary information of log
type LogData struct {
	EventLocation string    // the function where it throw the log
	Message       string    // log message string
	Level         EventType // denote the log is info, warning, error, etc.
}

func GetLogData(location, message string, level EventType) LogData {
	return LogData{
		EventLocation: location,
		Message:       message,
		Level:         level,
	}
}

// Declare variables to store log messages as new Events
const (
	// Warning
	DuplicateRzKeyMessage = "In mapping table, duplicate odiall rz key (%s), market id (%s)"
)

// Declare variables to store log messages as new Events
var (
	// Warning
	duplicateRzKey = Event{Warning, DuplicateRzKeyMessage}
)

func DuplicateRzKeyFormatString(rzKey, marketId string) string {
	return fmt.Sprintf(duplicateRzKey.message, rzKey, marketId)
}

func (l *StandardLogger) logInfo(logData LogData) {
	l.Infof(logData.Message)
}

func (l *StandardLogger) logWarn(logData LogData) {
	l.Warnf(logData.Message)
}

func (l *StandardLogger) logError(logData LogData) {
	l.Errorf(fmt.Sprintf("Error message: %s.\nError location: %s.", logData.Message, logData.EventLocation))
}
