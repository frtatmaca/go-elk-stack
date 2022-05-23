package main

import (
	"encoding/json"
	"github.com/go-elk/core/logger"
	"log"
	"net/http"
)

var (
	fileLogger     logger.LogClient = logger.NewLogClient(logger.NewLogger())
	rabbitmqLogger logger.LogClient = logger.NewLogClient(logger.NewLoggerRabbitMQ())
)

func init() {
}

func main() {
	http.Handle("/writinglogtofile", http.HandlerFunc(WritingLogToFile))
	http.Handle("/writinglogtorabbitmq", http.HandlerFunc(WritingLogToRabbitMQ))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func WritingLogToFile(w http.ResponseWriter, r *http.Request) {
	fileLogger.Info("Log is sent to FileBeat")
	keysJson, err := json.Marshal("Message is sent to FileBeat.")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(keysJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func WritingLogToRabbitMQ(w http.ResponseWriter, r *http.Request) {
	rabbitmqLogger.Info("Log is sent to RabbitMQ")

	keysJson, err := json.Marshal("Message is sent to RabbitMQ.")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(keysJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
