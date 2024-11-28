FROM golang:1.23

WORKDIR /todo-app/api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-app/api/bin/main ./cmd/server

EXPOSE 8080

CMD ["/todo-app/api/bin/main"]