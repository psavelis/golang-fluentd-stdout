package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Starting the server.")

	healthzFunc := http.HandlerFunc(Healthz)
	http.Handle("/healthz",LoggingMiddleware(healthzFunc))

	http.ListenAndServe(":6001", nil)
}

// Healthz is
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
