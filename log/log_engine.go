package logwrapper

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/logging"
)

func LogEngine(logChannel chan LogData) {
	l := NewLogger()
	for {
		logData := <-logChannel
		switch logData.Level {
		case Info:
			l.logInfo(logData)
		case Warning:
			l.logWarn(logData)
		case Error:
			l.logError(logData)
		}
	}
}

func LogGCloudEngine(logChannel chan LogData) {
	l := NewLogger()
	ctx := context.Background()
	projectID := os.Getenv("GCLOUD_PROJDCT_ID")
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	logName := os.Getenv("LOG_TITLE")
	logger := client.Logger(logName)
	defer func() {
		err := client.Close()
		if err != nil {
			l.logError(GetLogData("LogGCloudEngine", err.Error(), Error))
			logger.Log(logging.Entry{Payload: fmt.Sprintf("Error message: %s.\nError location: %s.", err.Error(), "LogGCloudEngine"), Severity: logging.Error})
		}
	}()
	for {
		logData := <-logChannel
		switch logData.Level {
		case Info:
			l.logInfo(logData)
			logger.Log(logging.Entry{Payload: fmt.Sprintf("Info message: %s.\nInfo location: %s.", logData.Message, logData.EventLocation), Severity: logging.Info})
		case Warning:
			l.logWarn(logData)
			logger.Log(logging.Entry{Payload: logData.Message, Severity: logging.Warning})
		case Error:
			l.logError(logData)
			logger.Log(logging.Entry{Payload: fmt.Sprintf("Error message: %s.\nError location: %s.", logData.Message, logData.EventLocation), Severity: logging.Error})
		}
	}
}
