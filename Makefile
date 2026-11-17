.PHONY:build run
build:
	@go build  -o bin/api ./cmd/api 
run: build
	@./bin/api
# Production build
pbuild:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o bin/api ./cmd/api
prun: pbuild
	@./bin/api