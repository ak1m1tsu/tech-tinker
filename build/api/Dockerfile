FROM golang:alpine AS builder
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -tags musl -v -o ./app ./cmd/api/main.go

FROM alpine:latest as runtime
RUN apk --no-cache add ca-certificates
COPY --from=builder /app /app
COPY --from=builder /configs/api.yaml /etc/tech-tinker/config.yaml
COPY --from=builder /build/certificates/public.pem /etc/tech-tinker/public.pem
ENTRYPOINT /app
