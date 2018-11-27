all: test

generate:
	easyjson ./auth/authManager.go

test:
	mkdir _test
	go test -v -race ./...
	go test -v -covermode=count -coverprofile=coverage.out ./...
