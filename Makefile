test:
	go test $(shell go list ./... | grep -v /vendor/)

run:
	go run main.go