APP_NAME=hoge

.PHONY: build
build: clean main.go
	env GOOS=linux go build -o bin/$(APP_NAME)

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	gofmt -w -s ./

.PHONY: lint
lint:
	golint --set_exit_status ./...
