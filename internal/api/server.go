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

// Чекпоинт 2: после добавления POST /items в спеку и go generate ./...
// компилятор попросит реализовать здесь новый метод интерфейса.
