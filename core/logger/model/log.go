package model

import "time"

type Log struct {
	UserName      string    `json:"x-UserName"`
	CorrelationId string    `json:"x-CorrelationId"`
	Msg           string    `json:"msg"`
	Message       error     `json:"message"`
	Type          string    `json:"type"`
	AppName       string    `json:"appName"`
	IpAddress     string    `json:"ipAddress"`
	LogTime       time.Time `json:"logTime"`
}

const (
	Error   = "Error"
	Warning = "Warning"
	Info    = "Info"
)
