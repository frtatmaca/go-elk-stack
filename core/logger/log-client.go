package logger

import (
	"github.com/go-elk/core/logger/model"
	"github.com/google/uuid"
	"log"
	"net"
	"time"
)

type LogClient struct {
	logger ILogger
}

func NewLogClient(logger ILogger) LogClient {
	return LogClient{logger: logger}
}

func (s LogClient) Info(msg string) {
	log := createLog(msg, model.Info)
	s.logger.Write(log)
}

func (s LogClient) Warning(msg string) {
	log := createLog(msg, model.Warning)
	s.logger.Write(log)
}

func (s LogClient) Error(msg string, err error) {
	log := createLog(msg, model.Error)
	log.Message = err
	s.logger.Write(log)
}

func createLog(msg string, logType string) model.Log {
	log := model.Log{}
	log.AppName = "Application 1"
	log.IpAddress = getOutboundIP().String()
	log.CorrelationId = uuid.New().String()
	log.Msg = msg
	log.Type = logType
	log.LogTime = time.Now()
	log.UserName = "frtatmaca"

	return log
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
