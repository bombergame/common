all: test

generate:
	go generate ./...

test:
	go test -v -race ./...
	go test -v -covermode=count -coverprofile=coverage.out ./...

clean:
	rm -rf ./coverage.out
