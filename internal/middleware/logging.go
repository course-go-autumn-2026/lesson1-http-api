package middleware

import (
	"log"
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

// Logging пишет в лог метод, путь и статус каждого ответа.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		log.Printf("%s %s -> %d", r.Method, r.URL.Path, rec.status)
	})
}
