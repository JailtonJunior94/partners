FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.work .
COPY address/go.mod address/go.sum ./address/
COPY proto/go.mod proto/go.sum ./proto/
COPY partner/go.mod partner/go.sum ./partner/

WORKDIR /app/address
RUN go mod download

WORKDIR /app
COPY proto/ ./proto/
COPY address/ ./address/

WORKDIR /app/address
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/address ./cmd/

FROM alpine:3.20
WORKDIR /app
RUN apk --no-cache add tzdata
ENV TZ=America/Sao_Paulo
COPY --from=builder /app/bin/address .
COPY --from=builder /app/address/settings.json .
EXPOSE 80
EXPOSE 9000
CMD ["/app/address"]
