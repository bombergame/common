all: test

test:
	go test -v -race ./...
	go test -v -covermode=count -coverprofile=coverage.out ./...

clean:
	rm -f coverage.out
