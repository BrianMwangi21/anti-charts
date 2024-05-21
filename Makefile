# Go related variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main

# Docker related variables
DOCKER=docker
DOCKER_IMAGE=anticharts

all: clean build run

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	./$(BINARY_NAME)
	
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

docker-build:
	$(DOCKER) build -t $(DOCKER_IMAGE) .

docker-run:
	$(DOCKER) run --rm -it $(DOCKER_IMAGE)

docker-clean:
	$(DOCKER) rmi $(DOCKER_IMAGE)
