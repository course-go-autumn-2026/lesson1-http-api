# Fallback для чекпоинта 2

Если go generate не отработал (сеть, прокси) - не застревайте:

1. Сверьте свой api/openapi.yaml с openapi.yaml.example - там уже добавлен POST /items
2. Скопируйте gen.go.example на место internal/api/gen.go:

       cp fallback/gen.go.example internal/api/gen.go

3. Дальше по заданию: реализуйте CreateItem в internal/api/server.go
