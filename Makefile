test:
	go test -cover $(shell go list ./... | grep -v /vendor/)

vet:
	go vet $(go list ./... | grep -v /vendor/)

run:
	go run main.go
