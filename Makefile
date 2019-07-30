compile:
	go build

test:
	@go test ./tests

test-update-snapshot:
	@UPDATE_SNAPSHOTS=true go test ./tests


deps:
	dep ensure
