package logger

import (
	"log"
	"net/http"
)

type Logger interface {
	Warning(payload string)
	Info(payload string)
	Critical(payload string)
	Data(data interface{})
	LogTrace(severity Severity, spanID, trace string, data interface{})
	LogHTTP(severity Severity, req *http.Request, data interface{})
	ErrorWarning(err error)
	ErrorCritical(err error)
}

var logger Logger

func SetLogger(lgr Logger) {
	logger = lgr
}

func Warning(payload string) {
	if logger == nil {
		log.Println("logger not set. please call SetLogger() first to set logger.")
		return
	}

	logger.Warning(payload)
}

func Info(payload string) {
	if logger == nil {
		log.Println("logger not set. please call SetLogger() first to set logger.")
		return
	}

	logger.Info(payload)
}

func Critical(payload string) {
	if logger == nil {
		log.Println("logger not set. please call SetLogger() first to set logger.")
		return
	}

	logger.Critical(payload)
}

func Data(data interface{}) {
	if logger == nil {
		log.Println("logger not set. please call SetLogger() first to set logger.")
		return
	}

	logger.Data(data)
}

func LogTrace(severity Severity, spanID, trace string, data interface{}) {
	if logger == nil {
		log.Println("logger not set. please call SetLogger() first to set logger.")
		return
	}

	logger.LogTrace(severity, spanID, trace, data)
}

func LogHTTP(severity Severity, req *http.Request, data interface{}) {
	if logger == nil {
		log.Println("logger not set. please call SetLogger() first to set logger.")
		return
	}

	logger.LogHTTP(severity, req, data)
}

func ErrorWarning(err error) {
	if logger == nil {
		log.Println("logger not set. please call SetLogger() first to set logger.")
		return
	}

	logger.ErrorWarning(err)
}

func ErrorCritical(err error) {
	if logger == nil {
		log.Println("logger not set. please call SetLogger() first to set logger.")
		return
	}

	logger.ErrorCritical(err)
}
