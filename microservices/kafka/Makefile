build:
	@go build -o ./bin/app ./cmd/.
	@chmod +x ./bin/app

start: build
	@./bin/app

test-race:
	@go clean -testcache
	@go test -race -v ./..

kafka-up:
	@docker-compose -f docker-compose.yaml up -d

kafka-down:
	@docker-compose -f docker-compose.yaml down
