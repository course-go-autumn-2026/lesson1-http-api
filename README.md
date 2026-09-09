# Занятие 1. HTTP API - заготовка

Рабочий проект для практики: сервер на net/http + OpenAPI-спека с кодогенерацией.
Всё уже собирается и запускается - на паре будете его дописывать.

## Сетап (сделать ДО пары, 10 минут)

Нужны только git и Go 1.22 или новее. Проверка: `go version`.
Go нет - ставим с https://go.dev/dl/ (на Windows берите msi-установщик).

```bash
git clone https://github.com/course-go-autumn-2026/lesson1-http-api.git
cd lesson1-http-api
go run ./cmd/server
```

В соседнем терминале:

```bash
curl -i localhost:8080/ping
```

Получили `200` и `{"status":"ok"}` - вы готовы, пришлите плюс в чат курса.
Что-то не так - пишите в чат, разберем до пары.

Запасной вариант, если локально не завелось: кнопка Code -> Create codespace
на странице репозитория. Откроется VS Code в браузере с готовым окружением,
все команды те же.

## Что где лежит

```
cmd/server/main.go    точка входа: mux, ручной /ping (образец для чекпоинта 1)
internal/api/         gen.go - сгенерированный код (руками не правим),
                      server.go - реализация ручек из спеки
api/openapi.yaml      контракт API
api/cfg.yaml          конфиг генератора
fallback/             запасные файлы, если генерация не отработает
```

oapi-codegen ставить не нужно: `go generate ./...` сам скачает и запустит
нужную версию - она прибита в internal/api/generate.go.

## Как работаем на паре

Один проект на группу. Кто-то один шарит экран и печатает - это драйвер.
Драйвер меняется на каждом чекпоинте. Остальные роли: штурман читает задание
и диктует, тестировщик готовит curl-команды.

Шаги ниже описаны на нейтральном примере items. Ваша группа делает то же самое
со своей сущностью из scenarios/ - у группы N сценарий N.

После пары пройдите чекпоинты сами: заготовка и эталон остаются в репе.

## Чекпоинт 1: своя ручка

Допишите в cmd/server/main.go хендлер `GET /items/{id}`:
верните JSON с полями id и title. Образец - /ping, path-параметр
достается через `r.PathValue("id")`.

Проверка: `curl -i localhost:8080/items/42` -> 200 и JSON.

## Чекпоинт 2: спека и генерация

1. Добавьте в api/openapi.yaml ручку `POST /items`: тело - схема ItemCreate
   (уже лежит в components), ответ 201 со схемой Item.
   В спеке пока нет ни одного POST, поэтому вот каркас - подставьте свои
   имена и следите за отступами, yaml к ним безжалостен:

   ```yaml
     /items:
       post:
         operationId: createItem
         requestBody:
           required: true
           content:
             application/json:
               schema:
                 $ref: "#/components/schemas/ItemCreate"
         responses:
           "201":
             description: created
             content:
               application/json:
                 schema:
                   $ref: "#/components/schemas/Item"
   ```

2. `go generate ./...`
3. Соберите проект - компилятор скажет, какого метода не хватает
4. Реализуйте его в internal/api/server.go (образец - GetVersion)

Проверка:

```bash
curl -i -X POST localhost:8080/items \
     -H "Content-Type: application/json" \
     -d '{"title":"lamp"}'
```

Ответ 201 и созданный item. Генерация упала - см. fallback/README.md.

## Чекпоинт 3: проверка и проектирование

Часть 1: вставьте свою спеку в https://editor.swagger.io - получите живую доку,
потыкайте ручки оттуда.

Часть 2: без кода. Расставьте ручки для «сервиса объявлений» - создать,
посмотреть одно, список активных, скрыть. Для каждой операции: метод + путь +
статус ответа. Разберем вместе в конце пары.

## После пары

- Конспект лекции - в репозитории конспектов курса
- ДЗ со звездочкой: логирующий middleware (метод, URI, статус ответа).
  Подсказки - в конспекте, раздел про middleware

## Troubleshooting

- `bind: address already in use` - на 8080 уже что-то висит:
  `lsof -ti :8080 | xargs kill` (mac/linux) или смените порт в main.go
- Windows: команды curl работают в PowerShell и cmd начиная с Windows 10
- `go generate` качает долго - это один раз, дальше из кеша
