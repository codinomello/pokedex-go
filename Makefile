BINARY_NAME=pokedex

build:
	go build -o $(BINARY_NAME) && ./$(BINARY_NAME)

default: build