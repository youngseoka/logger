package stdout

import (
	"fmt"
	"github.com/youngseoka/logger"
	"net/http"
	"time"
)

type Logger struct {}

func New() *Logger {
	return &Logger{}
}

var _ logger.Logger = &Logger{}

const logFormat = "[%s] %s | %v\n"

const (
	info     = "INFO"
	warn     = "WARN"
	critical = "CRITICAL"
	infoData = "DATA"
)

func (l *Logger) Warning(payload string) {
	fmt.Printf(logFormat, warn, time.Now().Format(time.RFC3339), payload)
}

func (l *Logger) Info(payload string) {
	fmt.Printf(logFormat, info, time.Now().Format(time.RFC3339), payload)
}

func (l *Logger) Critical(payload string) {
	fmt.Printf(logFormat, critical, time.Now().Format(time.RFC3339), payload)
}

func (l *Logger) Data(data interface{}) {
	fmt.Printf(logFormat, infoData, time.Now().Format(time.RFC3339), data)
}

func (l *Logger) LogTrace(severity logger.Severity, spanID, trace string, data interface{}) {
	fmt.Printf(logFormat, severity.String(), time.Now().Format(time.RFC3339), data)
}

func (l *Logger) LogHTTP(severity logger.Severity, req *http.Request, data interface{}) {
	fmt.Printf(logFormat, severity.String(), time.Now().Format(time.RFC3339), data)
}

func (l *Logger) ErrorWarning(err error) {
	fmt.Printf(logFormat, warn, time.Now().Format(time.RFC3339), err)
}

func (l *Logger) ErrorCritical(err error) {
	fmt.Printf(logFormat, critical, time.Now().Format(time.RFC3339), err)
}
