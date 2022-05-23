package logger

import (
	"encoding/json"
	"github.com/go-elk/core/logger/model"
	"log"
	"os"
	"path"
	"time"
)

type Logger struct{}

var (
	InfoLogger *log.Logger
)

var file *os.File
var err error

func NewLogger() *Logger {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		_ = os.Mkdir("logs", os.ModeDir)
	}

	fileName := time.Now().Format("2006-01-02") + "-logs.log"
	path := path.Join("logs", fileName)

	file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	return &Logger{}
}

func (s *Logger) Write(model model.Log) {
	msg, _ := json.Marshal(model)
	InfoLogger = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger.Println(string(msg))
}
