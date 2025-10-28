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

	// admin server
	admin := http.NewServeMux()
	mux.Handle("/admin/", http.StripPrefix("/admin", admin))

	admin.HandleFunc("GET /healthz", RouteHealthZ)
	admin.HandleFunc("GET /metrics", cfg.RouteMetrics)
	admin.HandleFunc("POST /reset", cfg.RouteMetricsReset)

	// api server
	api := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", api))

	api.HandleFunc("POST /validate_chirp", RouteValidateChirp)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
