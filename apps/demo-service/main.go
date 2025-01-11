package main

import (
	"net/http"
	"time"
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "demo",
			Subsystem: "http",
			Name:      "request_duration_seconds",
			Help:      "HTTP request latency",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	httpErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "demo",
			Subsystem: "http",
			Name:      "errors_total",
			Help:      "Total HTTP errors",
		},
		[]string{"path"},
	)
)

func main() {
	prometheus.MustRegister(httpDuration, httpErrors)

	http.Handle("/", instrument("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	}))

	http.Handle("/slow", instrument("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(500+rand.Intn(1000)) * time.Millisecond)
		w.Write([]byte("slow response\n"))
	}))

	http.Handle("/error", instrument("/error", func(w http.ResponseWriter, r *http.Request) {
		httpErrors.WithLabelValues("/error").Inc()
		http.Error(w, "error", http.StatusInternalServerError)
	}))

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}

func instrument(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		duration := time.Since(start).Seconds()
		httpDuration.WithLabelValues(path).Observe(duration)
	}
}
