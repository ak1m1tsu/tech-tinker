.PHONY: build
build:
	go build -o bin/tech-tinker-api ./cmd/app/main.go

.PHONY: run
run: build
	./bin/tech-tinker-api
