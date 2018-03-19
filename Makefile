GOPATH=$(shell pwd)


setup:
	go get "github.com/aws/aws-lambda-go/events"
	go get "github.com/aws/aws-lambda-go/lambda"
	go get "github.com/pmcdowell-okta/oktajwt"

clean:
	rm -rf bin/*

build:
	rm -rf bin/*
	mkdir -p bin
	@GOPATH=$(GOPATH) GOOS=linux GOARCH=amd64 go build -o ./bin/authorizer authorizer.go
	@GOPATH=$(GOPATH) GOOS=linux GOARCH=amd64 go build -o ./bin/vehicles vehicles.go
	@GOPATH=$(GOPATH) GOOS=linux GOARCH=amd64 go build -o ./bin/bookings bookings.go


deploy:
	sls deploy

