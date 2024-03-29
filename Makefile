test:
	go test -race ./...

lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yaml