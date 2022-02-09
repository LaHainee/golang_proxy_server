lint:
	golangci-lint run -c golangci.yml ./...

run:
	docker-compose up --build -d