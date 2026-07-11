package routes

import (
	"net/http"

	"http-server-projeto-korp/internal/handlers"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Setup(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/projeto-korp", handlers.ProjetoHandler)
	mux.Handle("/metrics", promhttp.Handler())
}
