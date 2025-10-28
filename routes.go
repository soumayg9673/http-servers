package main

import (
	"fmt"
	"net/http"
)

func RouteHealthZ(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 OK"))
}

func (cfg *apiConfig) RouteMetrics(w http.ResponseWriter, r *http.Request) {
	w.Write(fmt.Appendf(nil, "Hits: %v", cfg.fileserverHits.Load()))
	w.WriteHeader(http.StatusOK)
}

func (cfg *apiConfig) RouteMetricsReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.Write(fmt.Appendf(nil, "Hits: %v", cfg.fileserverHits.Load()))
	w.WriteHeader(http.StatusOK)
}
