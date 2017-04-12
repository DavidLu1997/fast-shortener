test:
	go test -cover $(shell go list ./... | grep -v /vendor/)

benchmark:
	go test -bench=. $(shell go list ./... | grep -v /vendor/)

test-coverage:
	./scripts/test

vet:
	go vet $(shell go list ./... | grep -v /vendor/)

run:
	go run main.go

lint:
	golint $(shell go list ./... | grep -v /vendor/) | grep -v -E 'exported|comment'

lint-all:
	golint $(shell go list ./... | grep -v /vendor/)
