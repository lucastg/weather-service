FROM golang:1.23-alpine AS build

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun

FROM alpine:latest AS certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /app
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /app/cloudrun .
COPY --from=build /app/.env .
EXPOSE 8080

ENTRYPOINT ["./cloudrun"]
