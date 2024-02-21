FROM golang:alpine AS builder
RUN apk add alpine-sdk
RUN apk add --no-cache git
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -tags musl -v -o ./app ./cmd/auth/main.go

FROM alpine:latest as runtime
RUN apk --no-cache add ca-certificates
COPY --from=builder /app /app
COPY --from=builder /configs/auth-api.yaml /etc/auth-api/config.yaml
COPY --from=builder /build/certificates/private.pem /etc/auth-api/private.pem
ENTRYPOINT /app
