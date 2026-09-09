package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/course-go-autumn-2026/lesson1-http-api/internal/api"
)

func main() {
	mux := http.NewServeMux()

	// Готовый хендлер - образец: так регистрируется ручка.
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})


	// ═══ ШАГ 1. Своя ручка ═══════════════════════════════════════════
	// Добавьте здесь хендлер GET /items/{id}: верните JSON с полями
	// id и title. Образец - /ping выше.
	// Подсказка: path-параметр достается через r.PathValue("id").
	// Проверка: curl -i localhost:8080/items/42
	// ═════════════════════════════════════════════════════════════════

	// Сгенерированные ручки из api/openapi.yaml - монтируются на тот же mux.
	api.HandlerFromMux(&api.Server{}, mux)

	// ═══ ШАГ 2 (продолжение). Подключите middleware ══════════════════
	// Когда допишете Logging в internal/middleware/logging.go,
	// оберните mux: замените mux ниже на middleware.Logging(mux)
	// и добавьте импорт:
	//   "github.com/course-go-autumn-2026/lesson1-http-api/internal/middleware"
	// Проверка: дерните любую ручку - в терминале сервера появится строка лога.
	// ═════════════════════════════════════════════════════════════════
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
