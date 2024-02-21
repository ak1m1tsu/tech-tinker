build: build-api build-auth build-fixture

build-api:
	GOOS=linux go build -v -o ./bin/publicapi ./cmd/api/main.go

build-auth:
	GOOS=linux go build -v -o ./bin/auth ./cmd/auth/main.go

build-fixture:
	GOOS=linux go build -v -o ./bin/fixture ./scripts/fixture/main.go
