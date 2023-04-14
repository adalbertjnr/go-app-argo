package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Define a new counter metric to track requests to the "/metrics" endpoint.
	requestCount := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	})

	// Register the counter with the Prometheus collector.
	prometheus.MustRegister(requestCount)

	// Define a new HTTP handler to serve the "/metrics" endpoint.
	http.Handle("/metrics", promhttp.Handler())

	// Define a new HTTP handler to serve the root route.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Start the HTTP server on port 5000.
	fmt.Printf("Ouvindo na porta 5000!")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
