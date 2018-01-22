package main

import (
	"fmt"
	"net/http"

	logging "github.com/psavelis/golang-fluentd-stdout/middlewares"
)

func main() {

	healthzFunc := http.HandlerFunc(Healthz)

	http.Handle("/healthz", logging.FluentdMiddleware(healthzFunc))

	fmt.Println("Started.")
	http.ListenAndServe(":80", nil)
}

// Healthz is
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
