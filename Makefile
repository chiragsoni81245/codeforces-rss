build:
	@go build -o bin/codeforces-rss cmd/codeforces-rss/main.go

run: build
	@./bin/codeforces-rss start $(ARGS)

test:
	@go test ./... -v --race
