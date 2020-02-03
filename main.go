package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Create custom metric
var httpRequestsCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "http_requests",
	Help: "number of requests",
})

func main() {
	// Endpoint for prometheus
	http.Handle("/metrics", promhttp.Handler())

	// Demo endpoint
	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		defer httpRequestsCounter.Inc()
		resp := "Hello ..."
		w.Write([]byte(resp))
	})

	http.ListenAndServe(":8080", nil)
}
