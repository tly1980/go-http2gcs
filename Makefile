compile:
	go build

test:
	go test ./tests


deps:
	dep ensure
