package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/course-go-autumn-2026/lesson1-http-api/internal/api"
	"github.com/course-go-autumn-2026/lesson1-http-api/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Своя ручка на стандартном mux, без кодогенерации: path-параметр
	// достается через r.PathValue("id").
	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"id": id, "title": "item " + id})
	})

	// Сгенерированные ручки из api/openapi.yaml - монтируются на тот же mux.
	api.HandlerFromMux(&api.Server{}, mux)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.Logging(mux)))
}
