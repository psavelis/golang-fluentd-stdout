package logging

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// ActionLog is
type ActionLog struct {
	LogType   string        `json:"type"`
	Timestamp time.Time     `json:"@timestamp"`
	Message   string        `json:"message"`
	Method    string        `json:"method"`
	Pid       int           `json:"pid"`
	Req       *RequestData  `json:"req"`
	Res       *ResponseData `json:"res"`
	Tags      []string      `json:"tags"`
}

// RequestData is
type RequestData struct {
	Headers       *http.Header `json:"headers"`
	Method        string       `json:"method"`
	Referer       string       `json:"referer"`
	RemoteAddress string       `json:"remoteAddress"`
	URL           string       `json:"url"`
	UserAgent     string       `json:"userAgent"`
}

// ResponseData is
type ResponseData struct {
	ResponseTime int64 `json:"responseTime"`
	StatusCode   int   `json:"statusCode"`
}

// FluentdMiddleware is
func FluentdMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		msg := "Request log*"

		tags := []string{"golang-fluentd", "test"}

		entry := NewActionLog("response", tags, msg, r)

		lrw := NewLogResponseWriter(w)
		h.ServeHTTP(lrw, r)

		requestLength := time.Since(entry.Timestamp)

		entry.Res = &ResponseData{
			StatusCode:   lrw.statusCode,
			ResponseTime: requestLength.Nanoseconds(),
		}

		data, err := json.Marshal(entry)

		if err != nil {
			fmt.Println(fmt.Sprintf("serialization error:%v", err))
		}

		fmt.Println(fmt.Sprintf("%s\n", data))
	})
}

// LogResponseWriter is
type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewLogResponseWriter is
func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{w, http.StatusOK}
}

// WriteHeader is
func (lrw *LogResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// NewActionLog is
func NewActionLog(logType string, tags []string, message string, r *http.Request) *ActionLog {

	timestamp := time.Now()

	entry := &ActionLog{
		LogType:   logType,
		Timestamp: timestamp,
		Message:   message,
		Method:    r.Method,
		Pid:       os.Getpid(),
		Tags:      tags,
		Req: &RequestData{
			URL:           r.URL.String(),
			Method:        r.Method,
			Headers:       &r.Header,
			Referer:       r.Referer(),
			UserAgent:     r.UserAgent(),
			RemoteAddress: r.RemoteAddr,
		},
	}

	return entry
}
