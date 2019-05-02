.PHONY: build clean deploy

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/hello hello/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
