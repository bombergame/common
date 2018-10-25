all: test

test:
	mkdir test
	go test -v -race ./...
	go test -v -covermode=count -coverprofile=./test/coverage.out ./...

clean:
	rm -rf ./test
