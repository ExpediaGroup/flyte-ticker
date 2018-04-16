all: ci

test:
	go test ./...

clean:
	rm -f flyte-ticker

build:
	go build

ci: clean test build

docker-build:
	docker build --rm -t flyte-ticker:latest .

docker-run:
	docker run --rm --name flyte-ticker -it -e FLYTE_API=http://localhost:8080 -e LOGLEVEL=DEBUG flyte-ticker:latest

.PHONY: all test clean build ci docker-build docker-run
