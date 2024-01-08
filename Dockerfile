FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY ./ ./

RUN swag init -g delivery/ws/hub.go

RUN CGO_ENABLED=0 GOOS=linux go build -o /chat-servis main.go

FROM scratch

COPY --from=builder /chat-servis /chat-servis
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8081
EXPOSE 8082

ENTRYPOINT ["/chat-servis"]
