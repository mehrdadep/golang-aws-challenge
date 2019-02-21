.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/addModel addModel/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/addDevice addDevice/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/getDevice getDevice/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/getModel getModel/main.go
clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
