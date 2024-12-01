FROM golang:1.23

COPY .env /todo-app/api/.env

COPY api/go.mod /todo-app/api/go.mod
COPY api/go.sum /todo-app/api/go.sum

COPY api/ /todo-app/api/

WORKDIR /todo-app/api
RUN go mod download 
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/main ./cmd/server/main.go

EXPOSE 8080

CMD ["/todo-app/api/bin/main"]