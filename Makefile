all: test build image run

test:
	docker run --rm -it -v $(shell pwd):/go/src/docker-go -w /go/src/docker-go golang:alpine go get -u -v github.com/golang/dep/cmd/dep ; go get -v github.com/tebeka/go2xunit; dep ensure -v ; 2>&1 go test -v | go2xunit -output tests.xml

build:
	docker run --rm -it -v $(shell pwd):/go/src/docker-go -w /go/src/docker-go golang:alpine go get -u -v github.com/golang/dep/cmd/dep ; dep ensure -v ; env GOOS=linux go build -v

image:
	docker build --no-cache -t docker-go:latest .

run:
	docker run -p 8080:8080 --rm --name docker-go-inst docker-go

clean:
	rm -f docker-go
	rm -f tests.xml
	rm -rf vendor

.PHONY: run build image