package middleware

import (
	"net/http"
)

// statusRecorder - обертка над ResponseWriter, запоминает статус ответа.
// В http.ResponseWriter нет способа прочитать статус, поэтому подсовываем
// хендлеру свою обертку: встраивание + переопределение одного метода.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

// ═══ ШАГ 2. Логирующий middleware ═══════════════════════════════════════
// Сейчас Logging ничего не делает - просто пропускает запрос дальше.
// Замените тело функции: оберните next так, чтобы в лог попадали
// метод, путь и статус ответа.
//
// Подсказка:
//   rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
//   next.ServeHTTP(rec, r)
//   log.Printf("%s %s -> %d", r.Method, r.URL.Path, rec.status)
// (не забудьте импорт "log")
// ════════════════════════════════════════════════════════════════════════
func Logging(next http.Handler) http.Handler {
	return next
}
