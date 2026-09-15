package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// Демо graceful shutdown. Запустите сервер, дерните curl localhost:8070/slow
// и сразу нажмите Ctrl+C в терминале сервера: он не оборвет запрос, а
// дождется ответа и только потом завершится - в логе видно оба события.
//
// Без Shutdown Ctrl+C убивает процесс мгновенно, а клиент получает
// оборванное соединение вместо ответа.
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /slow", func(w http.ResponseWriter, r *http.Request) {
		log.Println("slow: начал обработку, сплю 5с")
		time.Sleep(5 * time.Second)
		log.Println("slow: закончил, отвечаю")
		w.Write([]byte("done\n"))
	})

	srv := &http.Server{Addr: ":8070", Handler: mux}

	// ListenAndServe блокирует, поэтому запускаем сервер в горутине,
	// а в main ждем сигнал остановки.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("listening on :8070")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done() // ждем SIGINT/SIGTERM (Ctrl+C)
	log.Println("получен сигнал остановки, дожидаюсь активных запросов")

	// Shutdown перестает принимать новые соединения и ждет завершения
	// текущих, но не дольше таймаута. Контекст для него создаем заново:
	// ctx выше уже отменен сигналом.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("shutdown: ", err)
	}
	log.Println("остановлен корректно")
}
