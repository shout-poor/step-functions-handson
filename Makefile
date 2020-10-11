.PHONY: build clean deploy

node_modules:
	npm install

build: node_modules
	export GO111MODULE=on
	export GOPRIVATE="github.com/shout-poor"
	export GONOSUMDB="github.com/shout-poor"
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

