all: test

check:
	gometalinter --enable-all --enable-gc --vendor --line-length=100 \
		--deadline=1000s --exclude=.*_easyjson\.go ./...

generate:
	go generate ./...

test:
	go test -v -race ./...
	go test -v -covermode=count -coverprofile=coverage.out ./...

clean:
	rm -rf ./coverage.out
