test:
	go test -cover $(shell go list ./... | grep -v /vendor/)

test-coverage:
	./scripts/test

vet:
	go vet $(shell go list ./... | grep -v /vendor/)

run:
	go run main.go
