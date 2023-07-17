package httpsniffer

import (
	"github.com/fatih/color"
	"log"
	"net/http"
	"os"
	"time"
)

type HTTPLogger struct {
	handler http.Handler
	logger  *log.Logger
}

func NewHTTPLogger(handler http.Handler) *HTTPLogger {
	return &HTTPLogger{
		handler: handler,
		logger:  log.New(os.Stdout, "", 0),
	}
}

func (l *HTTPLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	recorder := &responseRecorder{ResponseWriter: w}
	l.handler.ServeHTTP(recorder, r)

	status := recorder.status
	duration := time.Since(startTime)

	l.logRequest(r, status, duration)
}

func (l *HTTPLogger) logRequest(r *http.Request, status int, duration time.Duration) {
	method := r.Method
	url := r.URL.Path

	methodColor := getColorByMethod(method)
	pathColor := color.New(color.FgCyan).SprintFunc()

	l.logger.Printf("%s ▶ %s %s", time.Now().Format("15:04:05.000"), methodColor(method), pathColor(url))
}

func getColorByMethod(method string) func(a ...interface{}) string {
	switch method {
	case http.MethodGet:
		return color.New(color.FgGreen).SprintFunc()
	case http.MethodPost:
		return color.New(color.FgCyan).SprintFunc()
	case http.MethodPut:
		return color.New(color.FgYellow).SprintFunc()
	case http.MethodDelete:
		return color.New(color.FgRed).SprintFunc()
	case http.MethodHead:
		return color.New(color.FgMagenta).SprintFunc()
	default:
		return color.New(color.Reset).SprintFunc()
	}
}

type responseRecorder struct {
	http.ResponseWriter
	status int
}

func (r *responseRecorder) Write(data []byte) (int, error) {
	return r.ResponseWriter.Write(data)
}

func (r *responseRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}
