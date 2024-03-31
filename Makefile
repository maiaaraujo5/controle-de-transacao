test:
	go test -race ./...

lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yaml

docker-compose-up: docker-build
	cd build/docker && docker-compose up -d

docker-compose-down:
	cd build/docker && docker-compose down

docker-build:
	docker build -t maiaaraujo5/controle-de-transacao:latest -f ./build/docker/Dockerfile .

generate-swagger:
	swag init --dir ./cmd/,./app/server/rest --output ./docs