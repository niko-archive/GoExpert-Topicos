# generate makefile for golang project

.DEFAULT_GOAL := run


run:
	@echo "==> RUN.."
	@go run cmd/server/main.go

build:
	@echo "==> BUILD.."
	@go build -o bin/server cmd/server/main.go ldflags "-s -w"

testx:
	@echo "==> TEST.."
	@go test -v ./...

clean:
	@echo "==> CLEAN.."
	@rm -rf bin/*

