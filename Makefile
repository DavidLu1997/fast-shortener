test:
	go test -cover $(shell go list ./... | grep -v /vendor/)

run:
	go run main.go
