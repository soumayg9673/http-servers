package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func RouteHealthZ(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 OK"))
}

func (cfg *apiConfig) RouteMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<html>
	<body>
		<h1>Welcome, Chirpy Admin</h1>
		<p>Chirpy has been visited %d times!</p>
	</body>
	</html>`, cfg.fileserverHits.Load())
	w.WriteHeader(http.StatusOK)
}

func (cfg *apiConfig) RouteMetricsReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.Write(fmt.Appendf(nil, "Hits: %v", cfg.fileserverHits.Load()))
	w.WriteHeader(http.StatusOK)
}

func RouteValidateChirp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type req struct {
		Body string `json:"body"`
	}
	requestBody := req{}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if len(requestBody.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	s := strings.Split(requestBody.Body, " ")
	for _, rpl := range []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	} {
		for i := range s {
			if strings.ToLower(s[i]) == rpl {
				s[i] = strings.ReplaceAll(s[i], s[i], "****")
			}
		}
	}

	type res struct {
		CleanedBody string `json:"cleaned_body"`
	}

	responseBody := res{
		CleanedBody: strings.Join(s, ""),
	}

	respondWithJSON(w, http.StatusOK, responseBody)
}
