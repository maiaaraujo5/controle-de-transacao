test:
	go test -race ./...

lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yaml

docker-compose-up:
	cd build/docker && docker-compose up -d