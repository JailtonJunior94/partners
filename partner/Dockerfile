FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go clean --modcache
RUN GOOS=linux go build -o server partner/server.go

FROM alpine:3.13
WORKDIR /app

RUN apk --no-cache add tzdata

ENV TZ=America/Sao_Paulo

COPY --from=builder /app/server .
COPY ["partner/settings.Development.json","."]
COPY ["partner/settings.Production.json","."]
COPY ["partner/settings.Staging.json","."]

EXPOSE 80
EXPOSE 8000

CMD [ "/app/server" ]