package stackdriver

import (
	"cloud.google.com/go/errorreporting"
	"cloud.google.com/go/logging"
	"context"
	"fmt"
	"github.com/youngseoka/logger"
	"log"
	"net/http"
	"os"
)

type Logger struct {
	logger      *logging.Logger
	logClient   *logging.Client
	errorClient *errorreporting.Client
	deployMode  string
	labels      map[string]string
	projectID   string
}

var _ logger.Logger = &Logger{}

type logPayload struct {
	Data string
}

func New(ctx context.Context, name, version, projectID string) (*Logger, error) {
	labels := make(map[string]string)

	labels["version"] = version

	deployMode := os.Getenv("DEPLOY_MODE")
	if deployMode == "" {
		log.Println("deploy mode not set. consume as test")
		name += "-" + "dev"
		deployMode = "dev"
	}

	labels["deployMode"] = deployMode

	logClient, err := logging.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	logger := logClient.Logger(name)

	errClient, err := errorreporting.NewClient(ctx, projectID, errorreporting.Config{
		ServiceName:    name,
		ServiceVersion: version,
		OnError: func(err error) {
			log.Printf("error reporting failed: %v", err)
		},
	})

	if err != nil {
		return nil, err
	}

	return &Logger{
		errorClient: errClient,
		logClient:   logClient,
		logger:      logger,
		deployMode:  deployMode,
		labels:      labels,
		projectID:   projectID,
	}, nil
}

func (l *Logger) Warning(payload string) {
	l.logger.Log(logging.Entry{
		Severity: logging.Warning,
		Labels:   l.labels,
		Payload: logPayload{
			Data: payload,
		},
	})
}

func (l *Logger) Critical(payload string) {
	l.logger.Log(logging.Entry{
		Severity: logging.Critical,
		Labels:   l.labels,
		Payload: logPayload{
			Data: payload,
		},
	})
}

func (l *Logger) Info(payload string) {
	l.logger.Log(logging.Entry{
		Severity: logging.Info,
		Labels:   l.labels,
		Payload: logPayload{
			Data: payload,
		},
	})
}

func (l *Logger) Data(data interface{}) {
	l.logger.Log(logging.Entry{
		Severity: logging.Default,
		Labels:   l.labels,
		Payload:  data,
	})
}

func (l *Logger) LogTrace(severity logger.Severity, spanID, trace string, data interface{}) {
	l.logger.Log(logging.Entry{
		Severity: logging.Severity(severity),
		Labels:   l.labels,
		Payload:  data,
		SpanID:   spanID,
		Trace:    l.traceID(trace),
	})
}

// Raw log
func (l *Logger) LogHTTP(severity logger.Severity, req *http.Request, data interface{}) {
	l.logger.Log(logging.Entry{
		HTTPRequest: &logging.HTTPRequest{
			Request: req,
		},
		Severity: logging.Severity(severity),
		Labels:   l.labels,
		Payload:  data,
	})
}

// ErrorCritical function
func (l *Logger) ErrorCritical(err error) {
	l.logger.Log(logging.Entry{
		Severity: logging.Critical,
		Labels:   l.labels,
		Payload: logPayload{
			Data: err.Error(),
		},
	})

	l.errorClient.Report(errorreporting.Entry{
		Error: err,
	})
}

// ErrorWarning function
// log to logging entry BUT DOES NOT report to error report
func (l *Logger) ErrorWarning(err error) {
	l.logger.Log(logging.Entry{
		Severity: logging.Warning,
		Labels:   l.labels,
		Payload: logPayload{
			Data: err.Error(),
		},
	})
}

func (l *Logger) Close() error {
	err := l.logger.Flush()
	if err != nil {
		return err
	}

	l.errorClient.Flush()
	err = l.errorClient.Close()
	if err != nil {
		return err
	}

	return l.logClient.Close()
}

func (l *Logger) traceID(traceID string) string {
	return fmt.Sprintf("projects/%s/traces/%s", l.projectID, traceID)
}
