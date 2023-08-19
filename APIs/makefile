# generate makefile for golang project

.DEFAULT_GOAL := run


run:
	@echo "==> RUN.."
	@go run cmd/server/main.go

build:
	@echo "==> BUILD.."
	@go build -o bin/server cmd/server/main.go

test:
	@echo "==> TEST.."
	@go test -v ./...

exp:
	@echo "==> EXPORT.."
	@export $(cat .env | xargs)

clean:
	@echo "==> CLEAN.."
	@rm -rf bin/*

