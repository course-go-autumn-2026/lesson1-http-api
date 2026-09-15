package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// То же самое, что в cmd/server, но на chi вместо стандартного mux.
// Сравните с cmd/server/main.go: методы, {id} и http.Handler совпадают,
// разница - в группах маршрутов и подключении middleware.
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger) // на все маршруты сразу, без ручной обертки mux

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		})

		r.Get("/items/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := chi.URLParam(req, "id")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"id": id, "title": "item " + id})
		})
	})

	log.Println("listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
