all: test

test:
	mkdir _test
	go test -v -race ./...
	go test -v -covermode=count -coverprofile=./_test/coverage.out ./...

clean:
	rm -rf ./_test
