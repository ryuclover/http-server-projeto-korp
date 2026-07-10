package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of requests to /projeto-korp endpoint.",
    })
)

func init() {
    prometheus.MustRegister(requestCounter)
}

type Response struct {
    Nome    string `json:"nome"`
    Horario string `json:"horario"`
}

// healthHandler provides a simple health check.
func healthHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

// projetoHandler serves the original JSON and increments the request counter.
func projetoHandler(w http.ResponseWriter, r *http.Request) {
    requestCounter.Inc()
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    resp := Response{
        Nome:    "Projeto Korp",
        Horario: time.Now().UTC().Format(time.RFC3339),
    }
    json.NewEncoder(w).Encode(resp)
}

func main() {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/projeto-korp", projetoHandler)
    http.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)

    log.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
