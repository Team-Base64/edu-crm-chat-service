FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY ./ ./

RUN swag init -g src/delivery.go

RUN CGO_ENABLED=0 GOOS=linux go build -o /chat-servis main.go

FROM scratch

COPY --from=builder /chat-servis /chat-servis

EXPOSE 8081
EXPOSE 8082

ENTRYPOINT ["/chat-servis"]
