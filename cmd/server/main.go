package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/course-go-autumn-2026/lesson1-http-api/internal/api"
)

func main() {
	mux := http.NewServeMux()

	// Ручной хендлер - образец для чекпоинта 1.
	// Свою ручку GET /items/{id} добавляйте рядом, по этому образцу.
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Сгенерированные ручки из api/openapi.yaml - монтируются на тот же mux.
	api.HandlerFromMux(&api.Server{}, mux)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
