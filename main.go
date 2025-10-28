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

	// file server
	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))

	// api server
	api := http.NewServeMux()
	mux.Handle("/admin/", http.StripPrefix("/admin", api))

	api.HandleFunc("GET /healthz", RouteHealthZ)
	api.HandleFunc("GET /metrics", cfg.RouteMetrics)
	api.HandleFunc("POST /reset", cfg.RouteMetricsReset)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
