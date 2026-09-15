# Занятие 1. HTTP API

Рабочий проект: сервер на net/http + OpenAPI-спека с кодогенерацией + логирующий
middleware. Клонируется и сразу запускается - весь код уже на месте.

## Запуск

Нужны git и Go 1.22 или новее. Проверка: `go version`.
Go нет - ставим с https://go.dev/dl/ (на Windows берите msi-установщик).

```bash
git clone https://github.com/course-go-autumn-2026/lesson1-http-api.git
cd lesson1-http-api
go run ./cmd/server
```

В соседнем терминале:

```bash
curl -i localhost:8080/ping
curl -i localhost:8080/items/42
curl -i -X POST localhost:8080/items -H "Content-Type: application/json" -d '{"title":"lamp"}'
```

## Что где лежит

```
cmd/server/main.go             точка входа: mux, ручка /ping, ручка /items/{id}
cmd/graceful-shutdown/main.go  демо graceful shutdown, отдельный сервер на :8090
internal/middleware/           логирующий middleware (метод, путь, статус ответа)
internal/api/                  gen.go - сгенерированный код (руками не правим),
                               server.go - реализация ручек из спеки
api/openapi.yaml                контракт API
api/cfg.yaml                    конфиг генератора
```

oapi-codegen ставить не нужно: `go generate ./...` сам скачает и запустит
нужную версию - она прибита в internal/api/generate.go.

## Демо: graceful shutdown

```bash
go run ./cmd/graceful-shutdown
```

В соседнем терминале дерните долгую ручку и сразу нажмите Ctrl+C в терминале
сервера, пока curl еще ждет ответ:

```bash
curl -i localhost:8090/slow
```

Сервер не оборвет запрос - дождется ответа (5 секунд) и только потом
завершится, в логе будет видно оба события по порядку. Без Shutdown
(например, `kill -9`) соединение обрывается мгновенно, curl получает
пустой ответ вместо `done`.

## Финал пары: проектирование ручек

Без кода, по группам. Все работают с одним и тем же items - расставьте
недостающие операции: список объявлений, обновление и удаление.
Метод + путь + статус ответа для каждой. Спорные точки для обсуждения:
DELETE или PATCH для "скрыть"; что вернет пустой список - 200 или 404.

## После пары

Конспект лекции - в репозитории конспектов курса.

## Troubleshooting

- `bind: address already in use` - на 8080 уже что-то висит:
  `lsof -ti :8080 | xargs kill` (mac/linux) или смените порт в main.go
- Windows: curl работает в PowerShell и cmd начиная с Windows 10
