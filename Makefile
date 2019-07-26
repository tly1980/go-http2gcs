compile:
	go build

test:
	go test -v ./tests

deps:
	dep ensure
