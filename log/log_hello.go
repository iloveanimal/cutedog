package logwrapper

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/logging"
)

func GcloudLogHello(msg string) {
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
			log.Println("CloudLog Client Close Failed")
		}
	}()

	logger.Log(logging.Entry{Payload: msg, Severity: logging.Info})

	logger.Log(logging.Entry{Payload: msg, Severity: logging.Warning})

	logger.Log(logging.Entry{Payload: msg, Severity: logging.Error})

}
