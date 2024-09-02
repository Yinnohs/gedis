start:
	@go run ./main.go

build:
	@go build -o ./bin/gedis ./main.go

.Phony: start build
