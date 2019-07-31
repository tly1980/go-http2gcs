compile:
	go build

test:
	@go test -v  ./tests

test-update-snapshot:
	@UPDATE_SNAPSHOTS=true go test -v ./tests


deps:
	dep ensure
