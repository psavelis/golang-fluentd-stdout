package main

import (
	"os"
	"fmt"
  	"net/http"
  	"time"
  	"encoding/json"
)

// LogEntry is
type LogEntry struct {
	Log string `json:"log"`
	Stream string `json:"stream"`
	Time string `json:"time"`
}

// LoggingMiddleware is
func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		entry := LogEntry{
			Log:	fmt.Sprintf("[%s]%v=>%s", time.Now(), r.Method, r.URL),
			Stream:	"stdout",
			Time: 	fmt.Sprintf("%v", time.Now()),
		}
		
		j, err := json.Marshal(entry)

		if err != nil {
			fmt.Println("error:", err)
		}
	    
		os.Stdout.Write(j)

		h.ServeHTTP(w, r)
	})
}