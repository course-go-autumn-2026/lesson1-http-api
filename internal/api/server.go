package api

import (
	"encoding/json"
	"net/http"
)

// Server реализует ServerInterface - сгенерированный контракт из api/openapi.yaml.
// Каждой ручке из спеки соответствует метод. Добавили ручку в спеку,
// перегенерили - компилятор сам подскажет, какой метод нужно написать.
type Server struct{}

// GetVersion - сгенерированный образец: ручка описана в спеке,
// роутинг и типы сгенерированы, руками пишется только эта функция.
func (s *Server) GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Version{Version: "1.0.0"})
}

// ═══ ШАГ 4. Реализация нового метода ════════════════════════════════════
// После шага 3 (раскомментировали POST /items + go generate) компилятор
// потребует метод CreateItem. Раскомментируйте каркас ниже и заполните
// пропуск - ответ должен быть 201 с созданным Item.
//
// func (s *Server) CreateItem(w http.ResponseWriter, r *http.Request) {
// 	var req ItemCreate
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
//
// 	// ═══ ШАГ 5. Валидация ═══════════════════════════════════════════
// 	// REST-правило: невалидный вход отбиваем на границе со статусом 400.
// 	// Добавьте здесь проверку: пустой req.Title -> 400 и return.
// 	// ════════════════════════════════════════════════════════════════
//
// 	w.Header().Set("Content-Type", "application/json")
// 	// Пропуск: поставьте статус 201 (w.WriteHeader) и верните
// 	// json-ответ Item{Id: 1, Title: req.Title} (образец - GetVersion выше).
// }
//
// Проверка:
//   curl -i -X POST localhost:8080/items -H "Content-Type: application/json" -d '{"title":"lamp"}'
// ════════════════════════════════════════════════════════════════════════
