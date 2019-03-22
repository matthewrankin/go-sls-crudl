.PHONY: build check clean deploy
GOBUILD=env GOOS=linux go build -ldflags="-s -w" -o

help:
	@echo "You can perform the following:"
	@echo ""
	@echo "  build   Build the Go code into ./bin"
	@echo "  check   Format, lint, vet, and test Go code"
	@echo "  clean   Remove the bin folder"
	@echo "  deploy  Clean, build, and then deploy to AWS"
	@echo ""

build: clean
	export GO111MODULE=on
	$(GOBUILD) bin/delete functions/delete/main.go
	$(GOBUILD) bin/get functions/get/main.go
	$(GOBUILD) bin/list-by-year functions/list-by-year/main.go
	$(GOBUILD) bin/post functions/post/main.go
	$(GOBUILD) bin/put functions/put/main.go

check:
	@echo 'Formatting, linting, vetting, and testing Go code.'
	go fmt ./...
	golint ./...
	go vet ./...
	go test ./...

clean:
	rm -rf ./bin

deploy: build
	sls deploy --verbose
