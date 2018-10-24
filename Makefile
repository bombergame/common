all: test

test:
	glide install
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls
	go test -v -race ./...
	go test -v -covermode=count -coverprofile=coverage.out ./...

clean:
	rm -f coverage.out
