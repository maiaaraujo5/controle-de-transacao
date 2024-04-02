test:
	go test -race ./app/...

e2e_test:
	go test -race ./e2e_tests/...

lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yaml

docker-compose-up: docker-build
	cd build/docker && docker-compose up -d

docker-compose-up-dependencies:
	cd build/docker && docker-compose up postgres -d

docker-compose-down:
	cd build/docker && docker-compose down

docker-build:
	docker build -t maiaaraujo5/controle-de-transacao:latest -f ./build/docker/Dockerfile .

generate-swagger:
	swag init --dir ./cmd/,./app/server/rest --output ./docs

run-go: docker-compose-up-dependencies
	CONF=./configs/default.yaml go run cmd/main.go
