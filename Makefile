build:
	@go build -o ./bin/goshu

run: build
	@./bin/goshu

test:
	@go test -v ./...