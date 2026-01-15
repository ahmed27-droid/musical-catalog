

run:
	go run ./cmd/app/main.go

dev:
	air

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...


vet:
	go vet ./...


