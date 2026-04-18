FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.work .
COPY address/go.mod address/go.sum ./address/
COPY proto/go.mod proto/go.sum ./proto/
COPY partner/go.mod partner/go.sum ./partner/

WORKDIR /app/partner
RUN go mod download

WORKDIR /app
COPY proto/ ./proto/
COPY partner/ ./partner/

WORKDIR /app/partner
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server ./cmd/

FROM alpine:3.20
WORKDIR /app
RUN apk --no-cache add tzdata
ENV TZ=America/Sao_Paulo
COPY --from=builder /app/bin/server .
COPY --from=builder /app/partner/settings.Development.json .
COPY --from=builder /app/partner/settings.Production.json .
EXPOSE 80
EXPOSE 8000
CMD ["/app/server"]
