# Используем официальный образ Golang
FROM golang:1.23

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Копируем конфигурационные файлы
COPY configs ./configs

# Собираем приложение
RUN go build -o /productgatewayapi ./cmd/main.go

# Открываем порт для доступа к приложению
EXPOSE 80

# Запускаем приложение
CMD ["/productgatewayapi"]