# Запустить сервер

go run cmd/app/main.go

# Все тесты в проекте

go test ./...

# С подробным выводом

go test -v ./...

# С покрытием кода

go test -cover ./...

go test -v ./internal/entity/...

go test -v ./internal/service/...

go test -v ./internal/delivery/http/v1/...

go test -v ./internal/repository/...

go test -v ./internal/DTO/...

Докер
docker-compose up --build
docker-compose down -v 

