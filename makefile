GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=geoip
DOCKER_IMAGE_NAME=zhangmingkai4315/geoip-server

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v 

test: 
	$(GOTEST) -v ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME) server

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .
