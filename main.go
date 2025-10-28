package main

import (
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	cfg := &apiConfig{}
	mux := http.NewServeMux()

	h := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", cfg.middlewareMetricsInc(h))

	mux.HandleFunc("GET /healthz", RouteHealthZ)
	mux.HandleFunc("GET /metrics", cfg.RouteMetrics)
	mux.HandleFunc("POST /reset", cfg.RouteMetricsReset)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
