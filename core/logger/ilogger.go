package logger

import . "github.com/go-elk/core/logger/model"

type ILogger interface {
	Write(Log)
}
